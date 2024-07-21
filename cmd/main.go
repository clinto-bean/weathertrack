package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/clinto-bean/weathertrack/internal/api"
	middleware "github.com/clinto-bean/weathertrack/internal/middleware"
	godotenv "github.com/joho/godotenv"
)

type apiConfig struct {
	// DB *database.Queries
	instance string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load environment variables: %v", err.Error())
	}

	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/alerts", api.HandlerGetAlerts)

	CORSHandler := middleware.CORSHandler(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: CORSHandler,
	}
	log.Printf("Server running on port %v\n", port)
	log.Fatal(srv.ListenAndServe())
}
