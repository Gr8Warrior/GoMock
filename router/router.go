package router

import (
	"github.com/gorilla/mux"
	"github.com/gr8warrior/mongomock/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", controller.ServerHome).Methods("GET")
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.AddMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
