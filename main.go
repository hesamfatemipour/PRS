package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func response(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "data": map[string]string{"message": message}}
}

func main() {
	router := gin.Default()

	// gets products
	router.GET("/products/:id", func(c *gin.Context) {
		// TODO GET Product From DB
		c.JSON(http.StatusOK, response(200, ""))
	})

	// adds products
	router.POST("/product/add", func(c *gin.Context) {
		// TODO insert product into DB
		c.JSON(http.StatusOK, response(204, "product has bean created successfully"))
	})

	// removes products
	router.DELETE("/products/:id", func(c *gin.Context) {
		// TODO Delete products from DB
		c.JSON(http.StatusOK, response(204, "product has bean deleted successfully"))
	})

	// search products
	router.GET("/products/search", func(c *gin.Context) {
		// retrieve data from redis by title
		name := c.Param("tile")
		c.JSON(http.StatusOK, response(200, name))
	})

	router.Run(":8080")

}
