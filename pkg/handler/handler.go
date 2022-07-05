package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

func createSearchMoviesHandler(s movies.MovieSearcher, sort bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")

		countMovies, _ := s.SearchCount(searchQuery)

		var movies []movies.Movie
		if countMovies > 0 {
			mod := countMovies % 10
			pages := countMovies / 10
			if mod > 0 {
				pages = pages + 1
			}
			for i := 1; i <= pages; i++ {
				// call service to get movies
				var moviesByPage, _ = s.SearchMovies(searchQuery, i)
				movies = append(movies, moviesByPage...)
				fmt.Printf("request %d time\n", i)
			}
		}

		if sort {
			// Sort Movies
			sortMoviesByYearAndTitle(movies)
		}
		// generate response
		response := make(map[string]interface{})
		response["result"] = movies
		//if err != nil {
		//	response["error"] = err.Error()
		//}
		json.NewEncoder(w).Encode(response)
	}
}

// Sort array by Year and title when year is equals
func sortMoviesByYearAndTitle(movies []movies.Movie) {
	//Sort movies array
	sort.Slice(movies, func(i, j int) bool {
		if movies[i].Year == movies[j].Year {
			return movies[i].Title < movies[j].Title
		}
		return movies[i].Year < movies[j].Year
	})
}

// NewHandler returns a router with all endpoint handlers
func NewHandler(s movies.MovieSearcher) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/movies", createSearchMoviesHandler(s, false)).Methods("GET")
	router.HandleFunc("/movies-sorted", createSearchMoviesHandler(s, true)).Methods("GET")
	return router
}
