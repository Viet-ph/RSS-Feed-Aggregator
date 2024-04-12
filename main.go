package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/handler"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/middleware"
	"github.com/joho/godotenv"
)

func main() {
	filePathRoot := "./static"
	godotenv.Load()
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.Handle("/app/*", http.StripPrefix("/app/", http.FileServer(http.Dir(filePathRoot))))
	mux.HandleFunc("GET /v1/readiness", handler.Readiness)

	corsMux := middleware.MiddlewareCors(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filePathRoot, port)
	log.Fatal(server.ListenAndServe())
}
