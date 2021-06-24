package comet

//go:generate protoc --micro_out=. --go_out=. comet.proto
//go:generate bash -c "protodoc comet.proto > comet.md"
