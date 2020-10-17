package handlers

import (
	"commons/proto/auth"

	"github.com/adityak368/swissknife/logger"
	"github.com/micro/go-micro/v2"
)

func InitHandlers(service micro.Service) {
	auth.RegisterAuthHandler(service.Server(), new(Auth))
	logger.Info.Println("Initialized Microservice Handlers")
}
