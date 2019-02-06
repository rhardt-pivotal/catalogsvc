package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	stdopentracing "github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	zip    = flag.String("zipkin", os.Getenv("ZIPKIN"), "Zipkin address")
	//	port        = flag.String("port", os.Getenv("CATALOG_PORT"), "Port number on which the service should run")
	//	ip          = flag.String("ip", os.Getenv("CATALOG_IP"), "Preferred IP address to run the service on")
	serviceName = "catalog"
)

const (
	dbName         = "acmefit"
	collectionName = "catalog"
)

// GetEnv accepts the ENV as key and a default string
// If the lookup returns false then it uses the default string else it leverages the value set in ENV variable
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	logger.Info("Setting default values for ENV variable " + key)
	return fallback
}

// This initiates a new Logger and defines the format for logs
func initLogger(f *os.File) {

	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "",
		PrettyPrint:     true,
	})

	// Set output of logs to Stdout
	// Change to f for redirecting to file
	logger.SetOutput(os.Stdout)

}

// This handles initiation of "gin" router. It also defines routes to various APIs
// Env variable CATALOG_IP and CATALOG_PORT should be used to set IP and PORT.
// For example: export CATALOG_PORT=8087 will start the server on local IP at 0.0.0.0:8087
func handleRequest() {

	router := gin.Default()

	router.Static("/static/images", "./images")

	v1 := router.Group("/")
	{
		v1.GET("/products", GetProducts)
		v1.GET("/products/:id", GetProduct)
		v1.POST("/products", CreateProduct)
	}

	//flag.Parse()

	// Set default values if ENV variables are not set
	port := GetEnv("CATALOG_PORT", "8082")
	ip := GetEnv("CATALOG_HOST", "0.0.0.0")

	ipPort := ip + ":" + port

	logger.Infof("Starting catalog service at %s on %s", ip, port)

	router.Run(ipPort)
}

// This is the main function. It creates a logger file, along with sessions to DB and
// a collector for tracer
func main() {

	//create your file with desired read/write permissions
	f, err := os.OpenFile("log.info", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Could not open file ", err)
	} else {
		initLogger(f)
	}

	dbsession := ConnectDB(dbName, collectionName, logger)

	logger.Infof("Successfully connected to database %s", dbName)

	zipkinCollector, err := zipkintracer.NewHTTPCollector("http://0.0.0.0:9411/api/v1/spans")
	if err != nil {
		logger.Fatalf("unable to create Zipkin HTTP collector: %+v", err)
	}
	defer zipkinCollector.Close()

	zipkinRecorder := zipkintracer.NewRecorder(zipkinCollector, false, "0.0.0.0:8080", "catalog")
	zipkinTracer, err := zipkintracer.NewTracer(zipkinRecorder, zipkintracer.ClientServerSameSpan(true), zipkintracer.TraceID128Bit(true))
	if err != nil {
		logger.Fatalf("unable to create Zipkin tracer: %+v", err)
	}

	stdopentracing.SetGlobalTracer(zipkinTracer)

	handleRequest()

	CloseDB(dbsession, logger)

	// defer to close
	defer f.Close()
}
