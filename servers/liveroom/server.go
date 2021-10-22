package liveroom

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	iconf "github.com/nano-kit/goeasy/internal/config"
	"github.com/nano-kit/goeasy/internal/redir"
	"github.com/nats-io/nats.go"
)

const (
	ServiceName = "liveroom"
)

type Server struct {
	Namespace          string   `json:"namespace"`
	Production         bool     `json:"production"`
	LogOutputPaths     []string `json:"logging_output_paths"`
	RedisServerAddress string   `json:"redis_server_address"`
	NatsServerAddress  string   `json:"nats_server_address"`

	natsConn *nats.Conn `json:"-"`
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

	// connect to redis database
	redisDB := redis.NewClient(&redis.Options{
		Addr: s.RedisServerAddress,
	})
	if err := redisDB.Ping(context.Background()).Err(); err != nil {
		// log with warn level because it can be reconnected later
		log.Warnf("Ping redis: %v", err)
	}

	// connect to nats
	var err error
	s.natsConn, err = nats.Connect(s.NatsServerAddress,
		nats.Name(s.Name()),
		nats.MaxReconnects(-1),
		nats.RetryOnFailedConnect(true),
	)
	if err != nil {
		// log with warn level because it can be reconnected later
		log.Warnf("Connect nats: %v", err)
	}

	// initialize the micro service
	var srvOpts []micro.Option
	srvOpts = append(srvOpts,
		micro.Name(s.Name()),
		micro.BeforeStop(s.beforeStop),
	)
	service := micro.NewService(srvOpts...)

	RegisterDemoHandler(service.Server(), new(Demo))
	room := new(Room)
	room.redisDB = redisDB
	room.natsConn = s.natsConn
	room.Init()
	RegisterRoomHandler(service.Server(), room)
	micro.RegisterSubscriber(s.Namespace+".topic.user-activity", service.Server(), room.onUserActivity,
		server.SubscriberQueue(s.Name()))

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	// close redis database
	redisDB.Close()
}

func (s *Server) beforeStop() error {
	// release all blocking calls, such as Flush() and NextMsg()
	s.natsConn.Close()
	return nil
}
