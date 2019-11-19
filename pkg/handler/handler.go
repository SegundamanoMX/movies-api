package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

func createSearchMoviesHandler(s movies.MovieSearcher, sort bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")

		// call service to get movies
		movies, err := s.SearchMovies(searchQuery, sort)

		// generate response
		response := make(map[string]interface{})
		response["result"] = movies
		if err != nil {
			response["error"] = err.Error()
		}
		json.NewEncoder(w).Encode(response)
	}
}

// NewHandler returns a router with all endpoint handlers
func NewHandler(s movies.MovieSearcher) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/movies", createSearchMoviesHandler(s, false)).Methods("GET")
	router.HandleFunc("/movies-sorted", createSearchMoviesHandler(s, true)).Methods("GET")
	return router
}
