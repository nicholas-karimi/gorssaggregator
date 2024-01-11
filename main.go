package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/nicholas-karimi/gorssaggregator/internal/database"

	_ "github.com/lib/pq" // prepeding _ tells to to include the pacjage in the program even if its not being called directly
)

type APIConfig struct {
	DB *database.Queries
}

func main() {

	// Get the port from the environment
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// BD
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL not found in the environment")
	}

	// establish db connection
	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// conver the conn(sql.DB) to a query
	queries := database.New(conn)

	apiCfg := APIConfig{
		DB: queries,
	}

	// routers
	router := chi.NewRouter()

	// cors handler
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	// scope the handler to receive get request for checking server health/readiness

	// v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/api/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %s", portString)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
