package config

var (
	// Version
	Version = "1.0.0"

	//AppName -- Name of the application
	AppName = "EmailService"

	// AddressMicroservice for the microservice to bind to
	AddressMicroservice = "0.0.0.0:4001"

	//AddressMessageBroker for the broker
	AddressMessageBroker = "broker.gostarter.com:4002"

	// LogFileName is the logfilename
	LogFileName = AppName + ".log"

	// EmailHost Config for sending emails
	EmailHost = "smtp.gostarter.com"

	// EmailPort Config for sending emails
	EmailPort = 25

	// EmailUsername Config for sending emails
	EmailUsername = ""

	// EmailPassword Config for sending emails
	EmailPassword = ""

	// FromEmailID Config for sending emails
	FromEmailID = "no-reply@gostarter.com"

	// JaegerURL for Opentracing
	JaegerURL = "0.0.0.0:5775"
)
