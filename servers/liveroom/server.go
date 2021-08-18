package liveroom

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	iconf "github.com/nano-kit/goeasy/internal/config"
)

const (
	ServiceName = "liveroom"
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
		log.Warnf("Ping redis: %v", err)
	}
	defer redisDB.Close()

	// initialize the micro service
	var srvOpts []micro.Option
	srvOpts = append(srvOpts, micro.Name(s.Name()))
	service := micro.NewService(srvOpts...)

	RegisterDemoHandler(service.Server(), new(Demo))
	room := new(Room)
	room.redisDB = redisDB
	RegisterRoomHandler(service.Server(), room)
	micro.RegisterSubscriber(s.Namespace+".topic.user-activity", service.Server(), room.onUserActivity,
		server.SubscriberQueue(s.Name()))

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
