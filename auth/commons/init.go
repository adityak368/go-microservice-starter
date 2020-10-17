package commons

import (
	"auth/config"
	"auth/db"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"

	"github.com/adityak368/swissknife/localization/i18n"
	"github.com/adityak368/swissknife/logger"
)

// Initialize common app code between microservice and httpservice
func InitApp() {

	logger.InitConsoleLogger()
	InitEnv()
	db.MongoConnect()

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		logger.Info.Printf("Caught Signal : %+v\n", sig)
		CleanUpApp()
		os.Exit(0)
	}()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.Error.Fatal(err)
	}
	localizer := i18n.Localizer()
	localizer.LoadJSONLocalesFromFolder(path.Join(dir, "res", "locales"))
}

// CleanUpApp bookkeeping microservice and httpservice
func CleanUpApp() {
	db.MongoClose()
	logger.DestroyLogger()
}

func InitEnv() {
	connectionString, isConnectionURIPresent := os.LookupEnv("MONGO_DB_URI")
	if isConnectionURIPresent {
		logger.Info.Println("MONGO_DB_URI environment variable is set. Using MONGO_DB_URI from env")
		config.DBUrl = connectionString
		logger.Info.Println("Connecting to mongodb " + connectionString)
	}

	httpConnectionString, isConnectionURIPresent := os.LookupEnv("AUTH_HTTP_URI")
	if isConnectionURIPresent {
		logger.Info.Println("AUTH_HTTP_URI environment variable is set. Using AUTH_HTTP_URI from env")
		config.AddressHTTP = httpConnectionString
		logger.Info.Println("Running http server on " + httpConnectionString)
	}

	serviceConnectionString, isConnectionURIPresent := os.LookupEnv("AUTH_SERVICE_URI")
	if isConnectionURIPresent {
		logger.Info.Println("AUTH_SERVICE_URI environment variable is set. Using AUTH_SERVICE_URI from env")
		config.AddressMicroservice = serviceConnectionString
		logger.Info.Println("Running service on " + serviceConnectionString)
	}

	brokerConnectionString, isConnectionURIPresent := os.LookupEnv("BROKER_URI")
	if isConnectionURIPresent {
		logger.Info.Println("BROKER_URI environment variable is set. Using BROKER_URI from env")
		config.AddressMessageBroker = brokerConnectionString
		logger.Info.Println("Connecting to broker " + brokerConnectionString)
	}

	jaegerConnectionString, isConnectionURIPresent := os.LookupEnv("JAEGER_URI")
	if isConnectionURIPresent {
		logger.Info.Println("JAEGER_URI environment variable is set. Using JAEGER_URI from env")
		config.JaegerURL = jaegerConnectionString
		logger.Info.Println("Connecting to jaeger " + jaegerConnectionString)
	}

	userSecretString, isUserSecretPresent := os.LookupEnv("USER_SECRET")
	if isUserSecretPresent {
		logger.Info.Println("USER_SECRET environment variable is set. Using USER_SECRET from env")
		config.UserSecret = userSecretString
	}
}
