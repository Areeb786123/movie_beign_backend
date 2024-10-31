package router

import (
	"corses/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/allMovies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controllers.GetMoviesById).Methods("GET")
	router.HandleFunc("/api/postMovie", controllers.AddMovie).Methods("POST")
	router.HandleFunc("/api/updateMovie/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/api/deleteMovie/{id}", controllers.DeleteMovieById).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovies", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}
