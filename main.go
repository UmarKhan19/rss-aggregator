package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not set")
	} else {
		router := chi.NewRouter()

		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers

		}))

		v1Router := chi.NewRouter()
		v1Router.Get("/healthz", handlerReadiness)
		v1Router.Get("/err", handlerErr)

		router.Mount("/v1", v1Router)

		server := &http.Server{
			Handler: router,
			Addr:    ":" + port,
		}

		log.Printf("Server starting on port %v", port)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}

}
