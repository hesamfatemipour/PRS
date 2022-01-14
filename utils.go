package main

func response(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "data": map[string]string{"message": message}}
}
