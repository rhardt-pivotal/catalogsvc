package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var (
	logger *logrus.Logger
	// zip    = flag.String("zipkin", os.Getenv("ZIPKIN"), "Zipkin address")
	// //	port        = flag.String("port", os.Getenv("CATALOG_PORT"), "Port number on which the service should run")
	// //	ip          = flag.String("ip", os.Getenv("CATALOG_IP"), "Preferred IP address to run the service on")
	// serviceName = "catalog"
)

const (
	dbName         = "acmefit"
	collectionName = "catalog"
)

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &jaegercfg.Configuration{
		ServiceName: "acmeshop",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: "http://192.168.152.218:14268/api/traces",
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

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

	// jLogger := jaegerlog.StdLogger
	// jMetricsFactory := metrics.NullFactory

	// // Initialize tracer with a logger and a metrics factory
	// tracer, closer, err := cfg.NewTracer(
	// 	jaegercfg.Logger(jLogger),
	// 	jaegercfg.Metrics(jMetricsFactory),
	// )

	tracer, closer := initJaeger("acmeshop")

	stdopentracing.SetGlobalTracer(tracer)

	handleRequest()

	CloseDB(dbsession, logger)

	defer closer.Close()

	// defer to close
	defer f.Close()
}
