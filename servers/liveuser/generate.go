package liveuser

//go:generate protoc --micro_out=. --go_out=. liveuser.proto
//go:generate bash -c "protodoc liveuser.proto > liveuser.md"
