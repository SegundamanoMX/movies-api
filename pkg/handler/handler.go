package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
	"sort"
	//"fmt"
	"strconv"
)


func createSearchMoviesHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")

		// call service to get movies
		movies, err := s.SearchMovies(searchQuery,"1")

		// generate response
		response := make(map[string]interface{})
		response["result"] = movies
		if err != nil {
			response["error"] = err.Error()
		}
		json.NewEncoder(w).Encode(response)
	}
}

func createSearchMoviesSortHandler(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// get parameter from request
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")
		response := make(map[string]interface{})
		//newMovies :=  make([]*movies, 0)

		for i := 1; i <= 100 ; i++{

			pageReq := strconv.Itoa(i)
			//fmt.Println(pageReq)

		 	// call service to get movies
			movies, err := s.SearchMovies(searchQuery,pageReq) 
			
			if len(movies) > 0 {
  				// generate response	
  				response["result"] = movies
				if err != nil {
					response["error"] = err.Error()
				}else {
					sort.SliceStable(movies, func(i, j int) bool {
						if movies[i].Year == movies[j].Year{
							return movies[i].Title < movies[j].Title
						}
    						return movies[i].Year < movies[j].Year
						})
					json.NewEncoder(w).Encode(response)
				}
			}else {

				break
			}
			
		}
		//
	}
}



// NewHandler returns a router with all endpoint handlers
func NewHandler(s movies.MovieSearcher) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/movies", createSearchMoviesHandler(s)).Methods("GET")
	router.HandleFunc("/movies-sorted", createSearchMoviesSortHandler(s)).Methods("GET")
	return router
}
