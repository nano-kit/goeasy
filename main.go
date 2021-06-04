package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	log "github.com/micro/go-micro/v2/logger"
	signalutil "github.com/micro/go-micro/v2/util/signal"
	"github.com/nano-kit/goeasy/comet"
	"github.com/nano-kit/goeasy/gate"
	"github.com/nano-kit/goeasy/internal/reexec"
	"github.com/nano-kit/goeasy/servers/imchat"
)

type serverType int
type serverRecord struct {
	server
	serverType
	cmd *exec.Cmd
}

const namespace = "io.goeasy"
const (
	stSupervisor serverType = iota
	stGate
	stComet
	stIMChat
)

var (
	runAs serverType = stSupervisor

	servers = []serverRecord{
		{
			server: &gate.Server{
				Address:   ":8080",
				Namespace: namespace,
			},
			serverType: stGate,
		},
		{
			server: &comet.Server{
				Namespace: namespace,
			},
			serverType: stComet,
		},
		{
			server: &imchat.Server{
				Namespace: namespace,
			},
			serverType: stIMChat,
		},
	}
)

func init() {
	for _, x := range servers {
		serverType := x.serverType
		reexec.Register(x.Name(), func() { runAs = serverType })
	}
	reexec.Init()
}

func findServerRecord(st serverType) server {
	for _, x := range servers {
		if x.serverType == st {
			return x.server
		}
	}
	panic(fmt.Sprintf("unknown server type: %v", st))
}

func main() {
	if runAs != stSupervisor {
		findServerRecord(runAs).Run()
		return
	}

	log.Init(log.WithFields(map[string]interface{}{"service": "supervisor"}))
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
