module auth

go 1.14

//TODO: HACK UNTIL ETCD IS FIXED
replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible

replace commons => ../commons

require (
	commons v0.0.0-00010101000000-000000000000
	github.com/adityak368/swissknife/localization v0.0.0-20201017141410-95d62b8ed51b
	github.com/adityak368/swissknife/logger v0.0.0-20201017144903-f2d40cf64617
	github.com/adityak368/swissknife/middleware v0.0.0-20201017151948-7f76c60fd0ab
	github.com/adityak368/swissknife/objectstore v0.0.0-20201017145442-bcc18635b2fd
	github.com/adityak368/swissknife/response v0.0.0-20201017145442-bcc18635b2fd
	github.com/adityak368/swissknife/validation v0.0.0-20201017155821-35bd2a7f21bf
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/didip/tollbooth v4.0.2+incompatible // indirect
	github.com/labstack/echo/v4 v4.1.17
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.mongodb.org/mongo-driver v1.4.2
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee
)
