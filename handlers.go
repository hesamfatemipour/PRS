package main

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func AddProduct(c *gin.Context) {
	var req Product
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response(400, "wrong data format"))

	}

	client := RedisClient()
	doc := redisearch.NewDocument("doc1", 1.0)
	doc.Set("title", req.name).
		Set("body", req.attributes).
		Set("date", time.Now().Unix()).
		Set("id", uuid.New().String())

	if err := client.IndexOptions(redisearch.DefaultIndexingOptions, doc); err != nil {
		c.JSON(http.StatusInternalServerError, response(500, "could not insert product into db"))

	}
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
	title := c.Param("title")
	docs, _ := SearchRedis(title)
	c.JSON(http.StatusOK, generateDocumentResponse(docs))
}
