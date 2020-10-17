package commons

import (
	"email/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/adityak368/swissknife/logger"
)

// InitApp Intializes common app code between microservice and httpservice
func InitApp() {

	logger.InitConsoleLogger()
	InitEnv()
	mailer := Mailer()
	mailer.StartDaemon()

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		logger.Info.Printf("Caught Signal : %+v\n", sig)
		CleanUpApp()
		os.Exit(0)
	}()

}

// CleanUpApp does the bookkeeping
func CleanUpApp() {
	mailer.StopDaemon()
	logger.DestroyLogger()
}

// InitEnv loads the environment variables
func InitEnv() {

	serviceConnectionString, isConnectionURIPresent := os.LookupEnv("EMAIL_SERVICE_URI")
	if isConnectionURIPresent {
		logger.Info.Println("EMAIL_SERVICE_URI environment variable is set. Using AUTH_SERVICE_URI from env")
		config.AddressMicroservice = serviceConnectionString
		logger.Info.Println("Running service on " + serviceConnectionString)
	}

	jaegerConnectionString, isConnectionURIPresent := os.LookupEnv("JAEGER_URI")
	if isConnectionURIPresent {
		logger.Info.Println("JAEGER_URI environment variable is set. Using JAEGER_URI from env")
		config.JaegerURL = jaegerConnectionString
		logger.Info.Println("Connecting to jaeger " + jaegerConnectionString)
	}
}
