package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"strings"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/v2/auth"
	"github.com/nano-kit/goeasy/comet"
	iauth "github.com/nano-kit/goeasy/internal/auth"
)

var (
	jsonMarshaler = jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: false,
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

var (
	uid = flag.String("uid", "comet_tester", "user identity")
	rid = flag.String("rid", "", "room identity")
)

func main() {
	flag.Parse()

	// connect to the websocket host
	url := "ws://127.0.0.1:8080/comet/subscribe"
	dialer := ws.Dialer{}
	ctx := context.Background()
	log.Println("DIAL " + url)
	conn, _, _, err := dialer.Dial(ctx, url)
	if err != nil {
		panic(err)
	}

	// first message must be auth
	acc := &auth.Account{ID: *uid, Issuer: "io.goeasy"}
	token := iauth.AccountToToken(acc)
	uplink := &comet.Uplink{
		T: comet.MsgType_AUTH,
		Auth: &comet.Auth{
			Token: token,
		},
		Join: &comet.JoinRoom{
			Rid: *rid,
		},
	}
	jsonBytes, _ := jsonMarshal(uplink)
	jsonStr := string(jsonBytes)
	log.Println("TX " + jsonStr)
	if err := wsutil.WriteClientText(conn, jsonBytes); err != nil {
		panic(err)
	}

	// processing loop
	for {
		buf, err := wsutil.ReadServerText(conn)
		if err != nil {
			panic(err)
		}
		log.Println("RX " + string(buf))

		var downlink comet.Downlink
		if err := jsonUnmarshal(buf, &downlink); err != nil {
			panic(err)
		}
		if downlink.T == comet.MsgType_HB {
			// when receiving a downlink heartbeat, send an uplink heartbeat
			if err := wsutil.WriteClientText(conn, []byte("{}")); err != nil {
				panic(err)
			}
			continue
		}
		if strings.HasPrefix(downlink.Push.Evt, "#join:") {
			buf, _ := jsonMarshal(&comet.Uplink{
				T: comet.MsgType_JOIN,
				Join: &comet.JoinRoom{
					Rid: strings.TrimPrefix(downlink.Push.Evt, "#join:"),
				},
			})
			if err := wsutil.WriteClientText(conn, buf); err != nil {
				panic(err)
			}
		}
	}
}
