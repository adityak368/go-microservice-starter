module email

go 1.14

replace commons => ../commons

require (
	commons v0.0.0-00010101000000-000000000000
	github.com/adityak368/swissknife/email v0.0.0-20201017141410-95d62b8ed51b
	github.com/adityak368/swissknife/logger v0.0.0-20201017144903-f2d40cf64617
	github.com/adityak368/swissknife/middleware v0.0.0-20201017141410-95d62b8ed51b // indirect
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
)
