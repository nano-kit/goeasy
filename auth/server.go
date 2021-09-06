package auth

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	pb "github.com/micro/go-micro/v2/auth/service/proto"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"
	sqlStore "github.com/micro/go-micro/v2/store/sqlite"
	authHandler "github.com/nano-kit/goeasy/auth/handler/auth"
	rulesHandler "github.com/nano-kit/goeasy/auth/handler/rules"
	iconf "github.com/nano-kit/goeasy/internal/config"
	"github.com/nano-kit/goeasy/internal/redir"
)

const (
	ServiceName = "auth"
)

type Server struct {
	Namespace      string   `json:"namespace"`
	Production     bool     `json:"production"`
	LogOutputPaths []string `json:"logging_output_paths"`
}

func NewServer() *Server {
	s := &Server{}
	if err := iconf.LoadInitialConfigFromFile("serverinit.yaml", s); err != nil {
		panic(err)
	}
	return s
}

func (s *Server) Name() string {
	return s.Namespace + "." + ServiceName
}

func (s *Server) Run() {
	log.Init(
		log.WithFields(map[string]interface{}{"service": ServiceName}),
		log.SetOption("outputs", s.LogOutputPaths),
		log.SetOption("color", !s.Production),
	)
	redir.RedirectStdoutStderrToFile(ServiceName, s.Production)

	// set the auth namespace
	auth.DefaultAuth.Init(auth.Namespace(s.Namespace))

	// initialize the micro service
	var srvOpts []micro.Option
	srvOpts = append(srvOpts, micro.Name("go.micro.auth")) // use this name for `micro` cli
	service := micro.NewService(srvOpts...)

	// setup the handlers
	ruleH := &rulesHandler.Rules{}
	authH := &authHandler.Auth{}

	// set the handlers store
	st := sqlStore.NewStore(store.Database(s.Namespace), store.Table(ServiceName))
	authH.Init(auth.Store(st))
	ruleH.Init(auth.Store(st))

	// register handlers
	pb.RegisterAuthHandler(service.Server(), authH)
	pb.RegisterRulesHandler(service.Server(), ruleH)
	pb.RegisterAccountsHandler(service.Server(), authH)

	// output setup infos
	log.Infof("tokenProvider=%v, store=%v", authH.TokenProvider, authH.Options.Store)

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
