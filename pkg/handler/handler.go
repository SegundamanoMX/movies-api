package handler

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

func createSearchMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")

		// call service to get movies
		movies, err := s.SearchMovies(searchQuery)

		// generate response
		response := make(map[string]interface{})
		response["result"] = movies
		if err != nil {
			response["error"] = err.Error()
		}
		json.NewEncoder(w).Encode(response)
	}
}

func createSearchMoviesSortedHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")

		// call service to get movies
		movies, err := s.SearchMovies(searchQuery)

		sort.SliceStable(movies, func(i, j int) bool { return movies[i].Title < movies[j].Title })
		sort.SliceStable(movies, func(i, j int) bool { return movies[i].Year < movies[j].Year })

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
	router.HandleFunc("/movies", createSearchMoviesHandler(s)).Methods("GET")
	router.HandleFunc("/movies-sorted", createSearchMoviesSortedHandler(s)).Methods("GET")
	return router
}
