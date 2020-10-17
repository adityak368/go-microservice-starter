package config

var (
	// Version is the Service Version
	Version = "1.0.0"

	//AppName is the Name of the application
	AppName = "AuthService"

	// AddressHTTP for the HTTP server
	AddressHTTP = "0.0.0.0:4002"

	// AddressMicroservice for the microservice to bind to
	AddressMicroservice = "0.0.0.0:4003"

	// AddressMessageBroker is the address of the broker for to receive events
	AddressMessageBroker = "localhost:4004"

	// LogFileName name
	LogFileName = AppName + ".log"

	// S3AvatarBucket is the Amazon storage Bucket name
	S3AvatarBucket = "go-starter-bucket"

	//UserSecret Secret Key for User JWT
	UserSecret = ""

	//DBUrl MongoDB Url
	DBUrl = "mongodb://localhost:27017"

	//DbName -- MongoDB DB Name
	DbName = "Auth"

	//DbTimeout is the request timeout in seconds
	DbTimeout = 10

	// JaegerURL for Opentracing
	JaegerURL = "localhost:5775"
)
