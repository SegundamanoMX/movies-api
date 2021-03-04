package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

func createSearchMoviesHandler(s movies.MovieSearcher, sortAscending bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")

		// call service to get movies
		movies, err := s.SearchMovies(searchQuery)

		// ascending sort if true
		if sortAscending {
			sort.SliceStable(movies, func(i, j int) bool {
				if strings.Compare(movies[i].Year, movies[j].Year) == 0 {
					return movies[i].Title < movies[j].Title
				}

				return movies[i].Year < movies[j].Year
			})
		}

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
