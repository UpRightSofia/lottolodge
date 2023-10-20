package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/UpRightSofia/lottolodge/src/models/config"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
)

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load()

	dbConfig := config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	models.NewPostgresStore(dbConfig, func(store *models.PostgresStore) {

		migrationsDir := "./src/models/migrations"

		err := goose.Up(store.GetDB(), migrationsDir)
		if err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}

		httpMux := http.NewServeMux()

		httpMux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		log.Fatal(http.ListenAndServe(":80", httpMux))
	})
}
