package main

import "github.com/RediSearch/redisearch-go/redisearch"

func response(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "data": map[string]string{"message": message}}
}

func generateDocumentResponse(documents []redisearch.Document) map[string]interface{} {
	var response map[string]interface{}
	for _, document := range documents {
		response["id"] = document.Id
		response["data"] = document.Payload
	}
}
