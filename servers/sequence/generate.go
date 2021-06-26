package sequence

//go:generate protoc --micro_out=. --go_out=. sequence.proto
//go:generate bash -c "protodoc sequence.proto > sequence.md"
