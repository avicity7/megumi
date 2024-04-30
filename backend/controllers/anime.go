package controllers

import (
	"encoding/json"
	"fmt"
	"megumi/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAnime(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")
	result, err := services.GetAnime(title)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	parsed, _ := json.Marshal(result)
	w.Write(parsed)
}
