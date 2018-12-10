package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ishrivatsa/catalogservice/catalog"
	"github.com/opentracing/opentracing-go"
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

func main() {

	dbsession := catalog.ConnectDB("catalog", "products")

	log.Printf("Successfully connected to mongodb")

	handleRequest()

	catalog.CloseDB(dbsession)
}
