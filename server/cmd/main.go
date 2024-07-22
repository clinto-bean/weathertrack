package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/clinto-bean/weathertrack/internal/api"
	middleware "github.com/clinto-bean/weathertrack/internal/middleware"
	godotenv "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load environment variables: %v", err.Error())
	}

	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	CORSHandler := middleware.CORSHandler(mux)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: CORSHandler,
	}
	mux.HandleFunc("/v1/alerts/*", api.HandlerHandleAlertRequests)

	log.Printf("Server running on port %v\n", port)
	log.Fatal(srv.ListenAndServe())
}
