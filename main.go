package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", HomeHandler).Methods("GET")

	port, exists := os.LookupEnv("PORT")
	if exists == false {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Application is running on localhost:" + port)

	log.Fatal(srv.ListenAndServe())
}
