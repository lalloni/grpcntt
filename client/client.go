package client

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/apex/log"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/lalloni/grpcntt/common"
	"github.com/lalloni/grpcntt/rpc"
)

func Ping(address string, usetls bool, count, size uint64, sleep, timeout time.Duration) error {

	log.Infof("pinging %d times to server at %s", count, address)

	payload, err := common.RandomBytes(size)
	if err != nil {
		return errors.Wrap(err, "generating random payload")
	}

	opts := []grpc.DialOption{}

	if usetls {
		log.Info("tls enabled")
		cfg := &tls.Config{
			InsecureSkipVerify: true,
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(cfg)))
	} else {
		log.Info("tls disabled")
		opts = append(opts, grpc.WithInsecure())
	}

	ctx := context.Background()

	c, cancel := context.WithTimeout(ctx, timeout)
	conn, err := grpc.DialContext(c, address, opts...)
	cancel()
	if err != nil {
		return errors.Wrap(err, "creating grpc client")
	}

	service := rpc.NewServiceClient(conn)

	for i := uint64(1); i <= count; i++ {
		c, cancel := context.WithTimeout(ctx, timeout)
		res, err := service.SimplePing(c, &rpc.PingRequest{
			Sequence: i,
			Time:     common.ToRPCTime(time.Now()),
			Payload:  payload,
		})
		cancel()
		if err != nil {
			return errors.Wrap(err, "pinging")
		}
		log.Infof("received %d: %s", i, common.FormatPingResponse(res))
		time.Sleep(sleep)
	}

	if err := conn.Close(); err != nil {
		return errors.Wrap(err, "disconnecting")
	}

	return nil

}
