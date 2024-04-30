package main

import (
	"megumi/config"
	"megumi/routes"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://open-house-bus-tracker.vercel.app", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)
	config.Connect(r)

	routes.Anime(r)

	env := os.Getenv("ENV")
	if env == "PROD" {
		http.ListenAndServeTLS(":3000", "fullchain.pem", "privkey.pem", r)
	} else {
		http.ListenAndServe("127.0.0.1:3000", r)
	}
}
