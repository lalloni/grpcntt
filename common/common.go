package common

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lalloni/grpcntt/rpc"
	"github.com/pkg/errors"
)

const nano = 1000000000

func ToRPCTime(t time.Time) *rpc.Time {
	n := t.UnixNano()
	return &rpc.Time{
		Seconds: n / nano,
		Nanos:   n % nano,
	}
}

func ToSTDTime(t *rpc.Time) time.Time {
	return time.Unix(t.Seconds, t.Nanos)
}

func FormatPingRequest(req *rpc.PingRequest) string {
	return fmt.Sprintf("ping request %d at %s with payload %d bytes", req.Sequence, ToSTDTime(req.Time).Format(time.RFC3339Nano), len(req.Payload))
}

func FormatPingResponse(res *rpc.PingResponse) string {
	t := ToSTDTime(res.Time)
	d := t.Sub(ToSTDTime(res.Request.Time))
	return fmt.Sprintf("ping response at %s (Δ %s) from %s", t.Format(time.RFC3339Nano), d, FormatPingRequest(res.Request))
}

func RandomBytes(size uint64) ([]byte, error) {
	rand.Seed(time.Now().UnixNano())
	bs := make([]byte, size)
	_, err := rand.Read(bs)
	if err != nil {
		return nil, errors.Wrap(err, "generating random bytes")
	}
	return bs, nil
}
