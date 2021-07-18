package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/nano-kit/goeasy/internal/proto"
)

func main() {
	const name = "io.goeasy.cli.pub"
	// create a service
	service := micro.NewService(
		micro.Name(name),
	)

	event := micro.NewEvent("io.goeasy.service.comet", service.Client())
	message := &proto.Message{
		Server: name,
		Event:  "demo",
		Time:   time.Now().UnixNano(),
	}
	bytes, _ := json.Marshal(message)
	event.Publish(context.Background(), &proto.Event{
		Evt: string(bytes),
	})
	time.Sleep(time.Second)
}
