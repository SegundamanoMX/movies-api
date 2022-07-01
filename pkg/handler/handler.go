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
		pageNumber := queryParams.Get("p")

		// call service to get movies
		movies, err := s.SearchMovies(searchQuery, pageNumber)

		// generate response
		response := make(map[string]interface{})
		response["result"] = movies

		if err != nil {
			response["error"] = err.Error()
		}

		json.NewEncoder(w).Encode(response)
	}
}

func createSearchAndSortMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")
		pageNumber := queryParams.Get("p")

		// call service to get movies
		movies, err := s.SearchMovies(searchQuery, pageNumber)

		sort.Slice(movies, func(i, j int) bool {

			if movies[i].Year == movies[j].Year {
				return movies[i].Title < movies[j].Title
			}

			return movies[i].Year < movies[j].Year
		})

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
	router.HandleFunc("/movies-sorted", createSearchAndSortMoviesHandler(s)).Methods("GET")
	return router
}
