package catalog

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type Product struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Picture     string        `json:"imageUrl"`
	Price       float32       `json:"price"`
	Tags        []string      `json:"tag"`
}

// GetProducts returns a list of all products
func GetProducts(c *gin.Context) {
	var products []Product

	error := collection.Find(nil).All(&products)

	if error != nil {
		message := "Products " + error.Error()
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": products})

}

// GetProduct returns a single product based on id
func GetProduct(c *gin.Context) {
	var product Product

	productID := c.Param("id")

	error := collection.FindId(bson.ObjectIdHex(productID)).One(&product)

	if error != nil {
		message := "Product " + error.Error()
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": product})

}

// CreateProduct adds a new product item to the database
func CreateProduct(c *gin.Context) {
	var product Product

	error := c.ShouldBindJSON(&product)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Field Name(s)/ Value(s)"})
		return
	}

	product.ID = bson.NewObjectId()

	error = collection.Insert(&product)

	if error != nil {
		message := "Product " + error.Error()
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Product created successfully!", "resourceId": product})

}
