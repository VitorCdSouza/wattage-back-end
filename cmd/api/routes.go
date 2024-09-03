package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewMux()
	
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	
	
	return mux
}