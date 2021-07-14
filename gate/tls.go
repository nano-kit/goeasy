package gate

import (
	"crypto/tls"
	"net"

	"github.com/micro/go-micro/v2/api/server"
	"github.com/micro/go-micro/v2/api/server/acme/autocert"
	log "github.com/micro/go-micro/v2/logger"
	maddr "github.com/micro/go-micro/v2/util/addr"
	mls "github.com/micro/go-micro/v2/util/tls"
)

func (s *Server) configureTLS(opts []server.Option) []server.Option {
	if s.EnableTLS {
		hosts := s.hosts()

		log.Infof("generate a certificate with hosts: %v", hosts)
		cert, err := mls.Certificate(hosts...)
		if err != nil {
			log.Fatalf("can not generate the certificate: %v", err)
		}
		config := &tls.Config{
			Certificates: []tls.Certificate{cert},
			NextProtos:   []string{"h2", "http/1.1"},
		}

		opts = append(opts, server.EnableTLS(true))
		opts = append(opts, server.TLSConfig(config))
	} else if s.EnableACME {
		if s.Domain == "" {
			log.Fatalf("need to set the domain name if enabled ACME")
		}

		opts = append(opts, server.EnableACME(true))
		opts = append(opts, server.ACMEHosts(s.Domain))
		opts = append(opts, server.ACMEProvider(autocert.NewProvider()))
	}

	return opts
}

func (s *Server) hosts() []string {
	if s.Domain != "" {
		return []string{s.Domain}
	}

	hosts := []string{s.Address}

	// check if its a valid host:port
	if host, _, err := net.SplitHostPort(s.Address); err == nil {
		if len(host) == 0 {
			hosts = maddr.IPs()
		} else {
			hosts = []string{host}
		}
	}

	return hosts
}
