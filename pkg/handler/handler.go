package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

func createSearchMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// Get parameters from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")
		page := "1"

		// Call function to search and process results
		response := processApiResult(s, searchQuery, page)

		json.NewEncoder(w).Encode(response)
	}
}

func createSortedMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// Getting parameters from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")
		page := "1"

		// Call function to search and process results
		response := processApiResult(s, searchQuery, page)
		// Sort data from response["result"] structure of type Movies
		sort.Sort(movies.Movies(response["result"].(movies.Movies)))

		json.NewEncoder(w).Encode(response)
	}
}

func processApiResult(s movies.MovieSearcher, searchQuery string, page string) map[string]interface{} {
	// Searches movies, checks number of results and iteratively searches for the rest
	var pagesTotal int
	response := make(map[string]interface{})
	result, count, err := s.SearchMovies(searchQuery, page)
	if count > 10 {
		if count%10 == 0 {
			pagesTotal = count / 10
		} else {
			pagesTotal = (count / 10) + 1
		}
		for i := 2; i <= pagesTotal; i++ {
			tmpResult, _, err := s.SearchMovies(searchQuery, strconv.Itoa(i))
			for j := 0; j < len(tmpResult); j++ {
				result = append(result, tmpResult[j])
			}
			if err != nil {
				response["error"] = err.Error()
				return response
			}
		}
		response["result"] = result
	} else {
		// result, err := s.SearchMovies(searchQuery, page)
		if err != nil {
			response["error"] = err.Error()
			return response
		}
		return response
	}
	return response
}

// NewHandler returns a router with all endpoint handlers
func NewHandler(s movies.MovieSearcher) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/movies", createSearchMoviesHandler(s)).Methods("GET")
	router.HandleFunc("/movies-sorted", createSortedMoviesHandler(s)).Methods("GET")
	return router
}
