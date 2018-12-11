package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ishrivatsa/catalogservice/catalog"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

const (
	dbName         = "catalog"
	collectionName = "products"
)

var tracer opentracing.Tracer

func handleRequest() {

	router := gin.Default()

	router.Static("/static/images", "./images")

	v1 := router.Group("/")
	{
		v1.GET("/products", catalog.GetProducts)
		v1.GET("/products/:id", catalog.GetProduct)
		//v1.POST("/products", catalog.CreateProduct)
	}

	router.Run(":8889")
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

	dbsession := catalog.ConnectDB(dbName, collectionName, logger)

	logger.Infof("Successfully connected to database %s", dbName)

	handleRequest()

	catalog.CloseDB(dbsession, logger)

	// defer to close
	defer f.Close()
}
