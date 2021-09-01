package liveroom

import (
	context "context"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/metadata"
)

type Demo struct{}

func (d *Demo) Hello(ctx context.Context, req *HelloReq, res *HelloRes) error {
	enterTime := time.Now()
	acc, ok := auth.AccountFromContext(ctx)
	if ok {
		res.Account = append(res.Account,
			&KV{Key: "id", Value: acc.ID},
			&KV{Key: "type", Value: acc.Type},
			&KV{Key: "issuer", Value: acc.Issuer},
			&KV{Key: "metadata", Value: fmt.Sprint(acc.Metadata)},
			&KV{Key: "scopes", Value: fmt.Sprint(acc.Scopes)},
		)
	}
	if clientIP, ok := metadata.Get(ctx, "X-Real-Ip"); ok {
		res.Account = append(res.Account,
			&KV{Key: "ipaddr", Value: clientIP},
		)
	}
	if md, ok := metadata.FromContext(ctx); ok {
		for k, v := range md {
			res.Account = append(res.Account,
				&KV{Key: k, Value: v},
			)
		}
	}
	if req.Sleep > 0 {
		time.Sleep(time.Duration(req.Sleep) * time.Second)
	}
	res.Ack = req.Say
	res.Time = float32(time.Since(enterTime).Seconds())
	return nil
}
