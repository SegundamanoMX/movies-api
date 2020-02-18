package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)
const pageElements int = 10
func createSearchMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")
		// call service to get movies
		movies,_, err := s.SearchMovies(searchQuery,"1")
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
		movies, _total,err := s.SearchMovies(searchQuery,"1")
		maxPage := _total/pageElements
		if (_total%pageElements > 0) {
			maxPage = _total/pageElements + 1
		}

		//get all pages for movies
		i := 2
		for (i <= maxPage) {
			movies2,_total,_err := s.SearchMovies(searchQuery, strconv.Itoa(i))
			_total = _total
			err = _err
			for _, movie := range movies2 {
				movies = append(movies, movie)
			}
			i = i + 1
		}
		//sorts the movie results by year ascending
		sort.Slice(movies, func(i, j int) bool {
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
	router.HandleFunc("/movies-sorted", createSearchMoviesSortedHandler(s)).Methods("GET")
	return router
}
