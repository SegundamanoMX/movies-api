package handler

import (
	//"sort"
	//"strconv"
	"log"
	"encoding/json"
	"net/http"

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
		log.Print("############RESPBODY######")
		log.Print(response)

		json.NewEncoder(w).Encode(response)
	}
}

func createSearchMoviesSortedHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")

		// call service to get movies
		moviesr, err := s.SearchMoviesSorted(searchQuery)

		// generate response
		response := make(map[string]interface{})
		response["result"] = moviesr
		if err != nil {
			response["error"] = err.Error()
		}
		log.Print("############RESPBODY######")
		/*peliculas := response["result"]

		peliculassort := []movies.MovieI{}

		for _, pe := range moviesr {
			yeari, err := strconv.Atoi(pe.Year)
			if err != nil {
				response["error"] = err.Error()
			}
			p := movies.MovieI{pe.Title, yeari}
			peliculassort = append(peliculassort, p)
		}

		sort.Sort(ByYear(peliculassort))

		response["result"] = peliculassort

		log.Print(peliculas)
		log.Print(peliculassort)*/

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
