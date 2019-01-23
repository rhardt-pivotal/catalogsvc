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
	logger      *logrus.Logger
	zip         = flag.String("zipkin", os.Getenv("ZIPKIN"), "Zipkin address")
	serviceName = "catalog"
)

const (
	dbName         = "catalog"
	collectionName = "products"
)

func handleRequest() {

	router := gin.Default()

	router.Static("/static/images", "./images")

	v1 := router.Group("/")
	{
		v1.GET("/products", GetProducts)
		v1.GET("/products/:id", GetProduct)
		//v1.POST("/products", CreateProduct)
	}

	router.Run(":8080")
}

func initLogger(f *os.File) {

	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "",
		PrettyPrint:     true,
	})

	//set output of logs to f
	logger.SetOutput(f)

}

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
