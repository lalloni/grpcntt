package server

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/apex/log"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/lalloni/grpcntt/common"
	"github.com/lalloni/grpcntt/rpc"
)

func Serve(org string, hosts []string, address string, usetls bool) error {

	log.Infof("starting server listening at %s", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return errors.Wrapf(err, "trying to listen on address %s", address)
	}

	if usetls {
		key, cert, err := common.GenerateSelfSignedCertificate(org, hosts, "", 1024, time.Now(), 24*time.Hour, false)
		if err != nil {
			return errors.Wrap(err, "generating self signed TLS certificate")
		}
		pair, err := tls.X509KeyPair(cert, key)
		if err != nil {
			return errors.Wrap(err, "preparing tls x509 key pair")
		}
		listener = tls.NewListener(listener, &tls.Config{
			Certificates: []tls.Certificate{pair},
		})
		log.Info("tls enabled")
	}

	server := grpc.NewServer()
	rpc.RegisterServiceServer(server, &serviceImpl{})

	return server.Serve(listener)

}

type serviceImpl struct {
	received uint64
}

func (s *serviceImpl) SimplePing(ctx context.Context, req *rpc.PingRequest) (*rpc.PingResponse, error) {
	s.received++
	log.Infof("received %d: %s", s.received, common.FormatPingRequest(req))
	return &rpc.PingResponse{
		Time:    common.ToRPCTime(time.Now()),
		Request: req,
	}, nil
}
