package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

// createSearchMoviesHandler requests the movies and generate the response
func createSearchMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		movies, err := getMoviesByTheQueryParameter(req, s)
		generateResponse(w, movies, err)
	}
}

// createSearchAndSortMoviesHandler requests the movies, sort them and generate the response
func createSearchAndSortMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		movies, err := getMoviesByTheQueryParameter(req, s)
		// sort movies by year (asc) and title (asc)
		sort.Slice(movies, func(i, j int) bool {
			if movies[i].Year != movies[j].Year {
				return movies[i].Year < movies[j].Year
			}
			return movies[i].Title < movies[j].Title
		})
		generateResponse(w, movies, err)
	}
}

// getMoviesByTheQueryParameter gets the query parameter to invokes the searching of movies
func getMoviesByTheQueryParameter(req *http.Request, s movies.MovieSearcher) ([]movies.Movie, error) {
	// get parameter from request
	queryParams := req.URL.Query()
	searchQuery := queryParams.Get("q")

	// call service to get movies
	movies, err := s.SearchMovies(searchQuery)
	return movies, err
}

// generateResponse generates the response with the gotten movies
func generateResponse(w http.ResponseWriter, movies []movies.Movie, err error) {
	// generate response
	response := make(map[string]interface{})
	response["result"] = movies
	if err != nil {
		response["error"] = err.Error()
	}
	encodeErr := json.NewEncoder(w).Encode(response)
	if encodeErr != nil {
		log.Printf("Error at encoding the response")
	}
}

// NewHandler returns a router with all endpoint handlers
func NewHandler(s movies.MovieSearcher) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/movies", createSearchMoviesHandler(s)).Methods("GET")
	router.HandleFunc("/movies-sorted", createSearchAndSortMoviesHandler(s)).Methods("GET")
	return router
}
