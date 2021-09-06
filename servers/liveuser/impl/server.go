package impl

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	iconf "github.com/nano-kit/goeasy/internal/config"
	ipath "github.com/nano-kit/goeasy/internal/path"
	"github.com/nano-kit/goeasy/internal/redir"
	"github.com/nano-kit/goeasy/servers/liveuser"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

const (
	ServiceName = "liveuser"
)

type Server struct {
	Namespace          string   `json:"namespace"`
	Production         bool     `json:"production"`
	LogOutputPaths     []string `json:"logging_output_paths"`
	RedisServerAddress string   `json:"redis_server_address"`
}

func NewServer() *Server {
	s := &Server{}
	if err := iconf.LoadInitialConfigFromFile("serverinit.yaml", s); err != nil {
		panic(err)
	}
	return s
}

func (s *Server) Name() string {
	return s.Namespace + ".service." + ServiceName
}

func (s *Server) Run() {
	log.Init(
		log.WithFields(map[string]interface{}{"service": ServiceName}),
		log.SetOption("outputs", s.LogOutputPaths),
		log.SetOption("color", !s.Production),
	)
	redir.RedirectStdoutStderrToFile(ServiceName, s.Production)

	// connect to sql database
	dir := filepath.Join(ipath.HomeDir(), "."+ServiceName)
	fname := ServiceName + ".db"
	os.MkdirAll(dir, 0700)
	dbPath := filepath.Join(dir, fname)
	dbURI := "file:" + dbPath + "?cache=shared"
	sqldb, err := sql.Open(sqliteshim.ShimName, dbURI)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	if err := sqldb.Ping(); err != nil {
		log.Infof("Ping sql: %v", err)
	}
	sqlDB := bun.NewDB(sqldb, sqlitedialect.New())
	sqlDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose()))

	// connect to redis database
	redisDB := redis.NewClient(&redis.Options{
		Addr: s.RedisServerAddress,
	})
	if err := redisDB.Ping(context.Background()).Err(); err != nil {
		// log with info level because it can be reconnected later
		log.Infof("Ping redis: %v", err)
	}

	// initialize the micro service
	var srvOpts []micro.Option
	srvOpts = append(srvOpts, micro.Name(s.Name()))
	service := micro.NewService(srvOpts...)

	// register subscriber
	user := new(User)
	user.Init(sqlDB, redisDB)
	liveuser.RegisterUserHandler(service.Server(), user)
	wx := new(Wx)
	wx.Init(service, s.Namespace)
	liveuser.RegisterWxHandler(service.Server(), wx)
	micro.RegisterSubscriber(s.Namespace+".topic.user-activity", service.Server(), user.onUserActivity,
		server.SubscriberQueue(s.Name()))

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	// close redis database
	redisDB.Close()

	// close sql database
	sqlDB.Close()
}
