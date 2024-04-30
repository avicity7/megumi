package routes

import (
	"megumi/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Anime(r *chi.Mux) {
	r.Route("/anime", func(r chi.Router) {
		r.Get("/get-anime/{title}", http.HandlerFunc(controllers.GetAnime))
	})
}
