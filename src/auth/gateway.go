package auth

import (
	"net/http"
	"os"
)

func CheckAuth(r *http.Request) bool {
	serviceApiKey := os.Getenv("SERVICE-API-KEY")
	apiKey := r.Header.Get("API-KEY")
	if serviceApiKey == "" || apiKey == serviceApiKey {
		return true
	}
	return false
}
