package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/v2/auth"
	"github.com/nano-kit/goeasy/comet"
	iauth "github.com/nano-kit/goeasy/internal/auth"
	"github.com/nano-kit/goeasy/internal/json"
)

var (
	jsonMarshaler = jsonpb.Marshaler{
		OrigName: true,
	}
	jsonUnmarshaler = jsonpb.Unmarshaler{}
)

func jsonMarshal(m proto.Message) ([]byte, error) {
	b := new(bytes.Buffer)
	err := jsonMarshaler.Marshal(b, m)
	return b.Bytes(), err
}

func jsonUnmarshal(data []byte, m proto.Message) error {
	return jsonUnmarshaler.Unmarshal(bytes.NewReader(data), m)
}

func main() {
	url := "ws://127.0.0.1:8080/comet/subscribe"
	dialer := ws.Dialer{}
	ctx := context.Background()
	conn, _, _, err := dialer.Dial(ctx, url)
	if err != nil {
		panic(err)
	}

	acc := &auth.Account{ID: "comet_tester"}
	token := iauth.AccountToToken(acc)

	if err := wsutil.WriteClientText(conn, []byte(fmt.Sprintf(`{"token":"%s"}`, token))); err != nil {
		panic(err)
	}

	for {
		buf, err := wsutil.ReadServerText(conn)
		if err != nil {
			panic(err)
		}
		log.Println("RAW " + string(buf))

		var push comet.ServerPush
		if err := jsonUnmarshal(buf, &push); err != nil {
			panic(err)
		}
		if push.T == comet.ServerPush_HEARTBEAT {
			if err := wsutil.WriteClientText(conn, []byte("{}")); err != nil {
				panic(err)
			}
			continue
		}
		log.Printf("JSON %v", json.Stringify(push))
	}
}
