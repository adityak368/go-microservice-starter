package main

import (
	"auth/commons"
	"auth/config"
	"auth/microservice"
	"auth/restapi/local"

	"github.com/adityak368/swissknife/logger"
	"github.com/adityak368/swissknife/middleware/errorhandler"
	"github.com/adityak368/swissknife/middleware/localization"
	"github.com/adityak368/swissknife/middleware/ratelimiter"
	"github.com/adityak368/swissknife/middleware/tracing"
	"github.com/adityak368/swissknife/validation/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/broker/nats"
	"github.com/micro/go-micro/v2/client/grpc"
)

func main() {

	commons.InitApp()

	// Init Tracing
	tracer := tracing.Init(
		tracing.OpenTracingConfig{
			JaegerURL:   config.JaegerURL,
			ServiceName: config.AppName,
		},
	)

	// start go-micro
	service := micro.NewService(
		micro.Name(config.AppName+".client"),
		micro.Client(grpc.NewClient()),
		micro.Broker(
			nats.NewBroker(
				broker.Addrs(config.AddressMessageBroker),
			),
		),
	)

	// Initialise service
	service.Init()

	// Initialize administration service
	microservice.InitService(service, true)

	logger.Info.Println("Starting HTTP server...")

	e := echo.New()
	e.Validator = playground.New()
	// Root level middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}","method":"${method}","uri":"${uri}","status":${status},"latency_human":"${latency_human}"}` + "\n",
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(tracing.EchoTracingMiddleware(tracer, "http."+config.AppName))
	e.Use(localization.EchoLocalizer())
	e.Use(ratelimiter.RateLimitMiddleware())
	e.HTTPErrorHandler = errorhandler.EchoHTTPErrorHandlerMiddleware

	auth := e.Group("/auth")
	local.InitLocalSignIn(auth)
	e.Logger.Fatal(e.Start(config.AddressHTTP))
}
