package imchat

import (
	context "context"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

const (
	ServiceName = "imchat"
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

type Demo struct{}

func (d *Demo) Hello(ctx context.Context, req *HelloReq, res *HelloRes) error {
	log.Infof("Entering Demo.Hello with req=%v", req)
	res.Ack = "<ack> " + req.Say
	return nil
}
