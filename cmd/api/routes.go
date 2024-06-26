package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)
	mux.Post("/authenticate", app.authenticate)
	mux.Get("/logout", app.logout)
	mux.Get("/refresh", app.refreshToken)
	mux.Get("/movies", app.AllMovies)
	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequired)
		mux.Get("/movies", app.MovieCatalog)
	})

	return mux
}
