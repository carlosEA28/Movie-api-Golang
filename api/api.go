package api

import (
	"net/http"

	"github.com/carlosEA28/helpers"
	"github.com/carlosEA28/omdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(apiKey string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handleSearchMovie(apiKey))

	return r
}

func handleSearchMovie(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pega o query parametro de search(o nome do filme)
		search := r.URL.Query().Get("s")
		res, err := omdb.Search(apiKey, search)

		if err != nil {
			helpers.SendJSON(w, helpers.Response{Error: "Something went wrong"}, http.StatusBadGateway)
			return
		}

		helpers.SendJSON(w, helpers.Response{Data: res}, http.StatusOK)
	}
}
