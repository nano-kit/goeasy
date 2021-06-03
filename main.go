package main

import "github.com/nano-kit/goeasy/gate"

func main() {
	gateServer := gate.Server{
		Address:   ":8080",
		Namespace: "io.goeasy",
	}
	gateServer.Run()
}
