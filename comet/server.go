package comet

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	iconf "github.com/nano-kit/goeasy/internal/config"
)

const (
	ServiceName = "comet"
)

type Server struct {
	Namespace string `json:"namespace"`
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
	log.Init(log.WithFields(map[string]interface{}{"service": ServiceName}))

	// initialize the micro service
	var srvOpts []micro.Option
	srvOpts = append(srvOpts, micro.Name(s.Name()))
	service := micro.NewService(srvOpts...)

	comet := NewComet(s.Namespace)
	RegisterCometHandler(service.Server(), comet)
	micro.RegisterSubscriber(s.Name(), service.Server(), comet.onEvent)

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
