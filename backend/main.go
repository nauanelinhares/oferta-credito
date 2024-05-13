package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nauanelinhares/ofertacredito/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	// Define the port
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Define dbUrl
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	// open sql

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Failed to open a DB connection:", err)
	}

	if err != nil {
		log.Fatal("Failed to create a new DB connection:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	// Person Routes
	v1Router.Post("/person", apiCfg.handlerCreatePerson)
	v1Router.Get("/person", apiCfg.handlerGetPersons)
	v1Router.Get("/person/{id}", apiCfg.handlerGetPerson)
	v1Router.Delete("/person/{id}", apiCfg.handlerDeletePerson)
	v1Router.Put("/person/{id}", apiCfg.handlerUpdatePerson)

	// Credit Request Routes
	v1Router.Post("/credit-request", apiCfg.handlerCreateCreditRequest)
	v1Router.Get("/credit-request", apiCfg.handlerGetCreditRequests)
	v1Router.Get("/credit-request/{id}", apiCfg.handlerGetCreditRequest)
	v1Router.Delete("/credit-request/{id}", apiCfg.handlerDeleteCreditRequest)
	v1Router.Put("/credit-request/{id}", apiCfg.handlerUpdateCreditRequest)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	// couting seconds app is executing
	go func() {
		timeStart := time.Now()
		timer := 0.0
		for {

			timeNow := time.Now()
			timer = timeNow.Sub(timeStart).Seconds()
			if timer > 60 {
				fmt.Println("Updating due")
				apiCfg.DB.UpdateDue(context.Background())
				timeStart = time.Now()
			}
		}
	}()

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
