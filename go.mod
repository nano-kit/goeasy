module github.com/nano-kit/goeasy

go 1.13

require (
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.7.3
	github.com/micro/go-micro/v2 v2.9.1
)

// Use a production ready go-micro/v2 stable version maintained by nano-kit.
replace github.com/micro/go-micro/v2 => ../go-micro
