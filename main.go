package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/UpRightSofia/lottolodge/src/handlers/pool"
	"github.com/UpRightSofia/lottolodge/src/handlers/tickets"
	"github.com/UpRightSofia/lottolodge/src/handlers/winnings"
	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/UpRightSofia/lottolodge/src/models/config"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
	"github.com/rs/cors"
)

func main() {
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

		healthRouter := mux.NewRouter()
		appRouter := mux.NewRouter()

		healthRouter.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		_ = tickets.NewServer(*store, appRouter)
		_ = pool.NewServer(*store, appRouter)
		_ = winnings.NewServer(*store, appRouter)

		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			log.Fatal(http.ListenAndServe(":80", healthRouter))
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			appRouter := cors.New(cors.Options{
				AllowedOrigins: []string{"*"},
				AllowedHeaders: []string{"*"},
				AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
			}).Handler(appRouter)
			log.Fatal(http.ListenAndServe(":8080", appRouter))
			wg.Done()
		}()
		wg.Wait()
	})
}
