package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProduct(c *gin.Context) {
	// TODO insert product into DB
	c.JSON(http.StatusOK, response(204, "product has bean created successfully"))
}


func GetProduct(c *gin.Context) {
	// TODO GET Product From DB
	c.JSON(http.StatusOK, response(200, ""))
}


func DeleteProduct(c *gin.Context) {
	// TODO Delete products from DB
	c.JSON(http.StatusOK, response(204, "product has bean deleted successfully"))
}

func SearchProducts(c *gin.Context) {
	// retrieve data from redis by title
	name := c.Param("tile")
	c.JSON(http.StatusOK, response(200, name))
}