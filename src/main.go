package main

import (
	"app/src/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("☁️ (cloud-storage-proxy) started")
	http.HandleFunc("/getBlob", handlers.GetBlob)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("no port found defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
