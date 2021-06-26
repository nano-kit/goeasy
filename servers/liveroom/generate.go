package liveroom

//go:generate protoc --micro_out=. --go_out=. demo.proto liveroom.proto
//go:generate bash -c "protodoc liveroom.proto > liveroom.md"
