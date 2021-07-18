package proto

//go:generate protoc --go_out=. proto.proto
//go:generate bash -c "protodoc proto.proto > proto.md"
