package main

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"log"
)

func initRedisClient() *redisearch.Client {
	c := redisearch.NewClient(RedisAddr, "primary")

	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewNumericField("id")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true}))

	err := c.CreateIndex(sc)
	if err != nil {
		c.Drop()
		log.Fatal(err)
	}
	return c
}

func RedisClient() *redisearch.Client {
	return redisearch.NewClient(RedisAddr, "primary")
}

func SearchRedis(q string) ([]redisearch.Document, int) {
	client := RedisClient()
	// Searching with limit and sorting
	docs, total, err := client.Search(redisearch.NewQuery(q).SetReturnFields("title"))
	if err != nil {
		log.Fatal("could not query redis")
	}
	return docs, total
}
