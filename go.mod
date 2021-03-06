module github.com/nano-kit/goeasy

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/disintegration/imaging v1.6.2
	github.com/go-redis/redis/v8 v8.11.3
	github.com/gobwas/ws v1.0.4
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/nats-io/nats.go v1.11.0
	github.com/prometheus/client_golang v1.11.0
	github.com/stretchr/testify v1.7.0
	github.com/uptrace/bun v1.0.8
	github.com/uptrace/bun/dialect/sqlitedialect v1.0.8
	github.com/uptrace/bun/driver/sqliteshim v1.0.8
	github.com/uptrace/bun/extra/bundebug v1.0.8
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
)

// Use a production ready go-micro/v2 stable version maintained by nano-kit.
replace github.com/micro/go-micro/v2 => ../go-micro
