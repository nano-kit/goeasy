package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	log "github.com/micro/go-micro/v2/logger"
	signalutil "github.com/micro/go-micro/v2/util/signal"
	"github.com/nano-kit/goeasy/auth"
	"github.com/nano-kit/goeasy/comet"
	"github.com/nano-kit/goeasy/gate"
	iconf "github.com/nano-kit/goeasy/internal/config"
	"github.com/nano-kit/goeasy/internal/reexec"
	"github.com/nano-kit/goeasy/servers/catalog"
	"github.com/nano-kit/goeasy/servers/liveroom"
	liveuser "github.com/nano-kit/goeasy/servers/liveuser/impl"
)

type serverName string
type serverRecord struct {
	server
	cmd *exec.Cmd
}

var (
	runAs serverName

	servers = []serverRecord{
		{server: auth.NewServer()},
		{server: gate.NewServer()},
		{server: comet.NewServer()},
		//{server: sequence.NewServer()},
		{server: liveroom.NewServer()},
		{server: liveuser.NewServer()},
		{server: catalog.NewServer()},
	}
)

func init() {
	for _, x := range servers {
		s := x.Name()
		reexec.Register(s, func() { runAs = serverName(s) })
	}
	reexec.Init()
}

func findServerRecord(s serverName) server {
	for _, x := range servers {
		if x.Name() == string(s) {
			return x.server
		}
	}
	panic(fmt.Sprintf("unknown server name: %v", s))
}

func main() {
	if runAs != "" {
		findServerRecord(runAs).Run()
		return
	}

	type Server struct {
		Production     bool     `json:"production"`
		LogOutputPaths []string `json:"logging_output_paths"`
	}
	s := &Server{}
	if err := iconf.LoadInitialConfigFromFile("serverinit.yaml", s); err != nil {
		panic(err)
	}
	log.Init(
		log.WithFields(map[string]interface{}{"service": "supervisor"}),
		log.SetOption("outputs", s.LogOutputPaths),
		log.SetOption("color", !s.Production),
	)
	log.Info("Start")
	term := make(chan os.Signal, 1)

	// start servers in reverse order
	// https://jiajunhuang.com/articles/2018_03_08-golang_fork.md.html
	for i := len(servers) - 1; i >= 0; i-- {
		server := servers[i]
		cmd := reexec.Command(server.Name())
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			log.Fatalf("failed to run command: %v", err)
			goto Quit
		} else {
			server.cmd = cmd
		}
	}

	signal.Notify(term, signalutil.Shutdown()...)
	// block until a signal is received.
	<-term

Quit:
	// stop servers in reverse order
	for i := len(servers) - 1; i >= 0; i-- {
		server := servers[i]
		if server.cmd != nil {
			server.cmd.Process.Signal(syscall.SIGTERM)
		}
	}
	for i := len(servers) - 1; i >= 0; i-- {
		server := servers[i]
		if server.cmd != nil {
			if err := server.cmd.Wait(); err != nil {
				log.Fatalf("failed to wait command: %v", err)
			}
		}
	}

	log.Info("Done")
}
