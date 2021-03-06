package handler

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

func generateResponse(m []movies.Movie, err error) map[string]interface{} {
	// Utility function to handle and generate response
	response := make(map[string]interface{})
	response["result"] = m
	if err != nil {
		response["error"] = err.Error()
	}
	return response
}

func createSearchMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")
		page := queryParams.Get("page")

		// call service to get movies
		movies, err := s.SearchMovies(searchQuery, page)

		// generate response
		response := make(map[string]interface{})
		response["result"] = movies
		if err != nil {
			response["error"] = err.Error()
		}
		json.NewEncoder(w).Encode(response)
	}
}

func createSortedMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// Getting parameters from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")
		page := queryParams.Get("page")

		// Calling search service
		movies_struct, err := s.SearchMovies(searchQuery, page)

		// Sorting data from struct
		sort.Sort(movies.Movies(movies_struct))

		// Generating and formatting response
		response := generateResponse(movies_struct, err)
		json.NewEncoder(w).Encode(response)
	}
}

// NewHandler returns a router with all endpoint handlers
func NewHandler(s movies.MovieSearcher) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/movies", createSearchMoviesHandler(s)).Methods("GET")
	router.HandleFunc("/movies-sorted", createSortedMoviesHandler(s)).Methods("GET")
	return router
}
