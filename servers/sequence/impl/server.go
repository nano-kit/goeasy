package impl

import (
	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	iconf "github.com/nano-kit/goeasy/internal/config"
	"github.com/nano-kit/goeasy/internal/redir"
	seq "github.com/nano-kit/goeasy/servers/sequence"
)

const (
	ServiceName = "sequence"
)

type Server struct {
	Namespace          string   `json:"namespace"`
	Production         bool     `json:"production"`
	LogOutputPaths     []string `json:"logging_output_paths"`
	RedisServerAddress string   `json:"redis_server_address"`

	redisDB *redis.Client
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

	// initialize server's dependent resources
	s.init()

	// initialize the micro service
	var srvOpts []micro.Option
	srvOpts = append(srvOpts, micro.Name(s.Name()))
	service := micro.NewService(srvOpts...)

	seq.RegisterSequenceHandler(service.Server(), &Sequence{s})

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) init() {
	// connect to the redis server
	s.redisDB = redis.NewClient(&redis.Options{
		Addr: s.RedisServerAddress,
	})
}
