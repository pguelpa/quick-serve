package main

import (
	"net/http"
	"os"
)

func main() {
	port := getConfigFromEnv("PORT", "8080")
	root := getConfigFromEnv("ROOT", ".")

	panic(http.ListenAndServe(":"+port, http.FileServer(http.Dir(root))))
}

func getConfigFromEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
