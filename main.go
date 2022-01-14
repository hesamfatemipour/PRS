package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(gin.Recovery())


	router.GET("/products/:id", GetProduct)
	router.POST("/product/add", AddProduct)
	router.DELETE("/products/:id", DeleteProduct)
	router.GET("/products/search", SearchProducts)

	router.Run(":8080")
}
