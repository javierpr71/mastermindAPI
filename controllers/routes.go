package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/javierpr71/mastermind/middlewares"
)

func (s *Server) initializeRoutes() *mux.Router {

	r := mux.NewRouter()

	// handler for documentation
	redocOpts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	docHandler := middleware.Redoc(redocOpts, nil)
	r.Handle("/docs", docHandler).Methods(http.MethodGet)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./"))).Methods(http.MethodGet)

	// Home Route
	r.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods(http.MethodGet)

	// Health
	//r.HandleFunc("/health", middlewares.SetMiddlewareJSON(s.Health)).Methods(http.MethodGet)

	// MasterMind
	r.HandleFunc("/newgame", middlewares.SetMiddlewareJSON(s.NewGame)).Methods(http.MethodPost)
	r.HandleFunc("/round", middlewares.SetMiddlewareJSON(s.Round)).Methods(http.MethodPost)
	r.HandleFunc("/status/{id}", middlewares.SetMiddlewareJSON(s.RoundStatus)).Methods(http.MethodGet)

	return r
}
