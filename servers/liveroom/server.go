package liveroom

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

const (
	ServiceName = "liveroom"
)

type Server struct {
	Namespace string
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

	RegisterDemoHandler(service.Server(), new(Demo))

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
