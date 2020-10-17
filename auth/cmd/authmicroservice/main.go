package main

import (
	"auth/commons"
	"auth/config"
	"auth/microservice"

	"github.com/adityak368/swissknife/logger"
	"github.com/adityak368/swissknife/middleware/tracing"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/broker/nats"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/server"
	grpcServer "github.com/micro/go-micro/v2/server/grpc"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
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

	// New Service
	service := micro.NewService(
		micro.Name(config.AppName),
		micro.Version("latest"),
		micro.Address(config.AddressMicroservice),
		micro.Client(grpc.NewClient()),
		micro.Server(
			grpcServer.NewServer(
				server.Name("micro."+config.AppName),
				server.Address(config.AddressMicroservice),
				server.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
				server.WrapSubscriber(opentracing.NewSubscriberWrapper(tracer)),
			),
		),
		micro.Broker(
			nats.NewBroker(
				broker.Addrs(config.AddressMessageBroker),
			),
		),
	)

	// Initialise service
	service.Init()

	// Initialize auth service
	microservice.InitService(service, false)

	// Run service
	if err := service.Run(); err != nil {
		logger.Error.Fatal(err)
	}
}
