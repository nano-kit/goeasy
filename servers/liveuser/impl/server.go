package impl

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	iconf "github.com/nano-kit/goeasy/internal/config"
	"github.com/nano-kit/goeasy/servers/liveuser"
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
	user.redisDB = redisDB
	liveuser.RegisterUserHandler(service.Server(), user)
	micro.RegisterSubscriber(s.Namespace+".topic.user-activity", service.Server(), user.onUserActivity,
		server.SubscriberQueue(s.Name()))

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	// close redis database
	redisDB.Close()
}
