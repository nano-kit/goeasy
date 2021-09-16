package catalog

//go:generate protoc --micro_out=. --go_out=. catalog.proto
//go:generate bash -c "protodoc catalog.proto > catalog.md"
