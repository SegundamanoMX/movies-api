package handler

import (
	"encoding/json"
	"fmt"
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
		page := queryParams.Get("page")
		if page == "" {
			page = "1"
		}

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

func CreateSearchMoviesHandlerSort(s movies.MovieSearcher) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		queryParams := req.URL.Query()
		searchQuery := queryParams.Get("q")
		sortQuery := queryParams.Get("sort")
		page := queryParams.Get("page")
		if page == "" {
			page = "1"
		}

		response := make(map[string]interface{})
		movies, err := s.SearchMovies(searchQuery, page)
		if err != nil {
			response["error"] = err.Error()
		}

		fmt.Println("***********************")
		fmt.Println(movies)

		if sortQuery != "" {
			sort.SliceStable(movies, func(i, j int) bool {
				result := false
				if movies[i].Year == movies[j].Year {
					listName := []string{movies[i].Title, movies[j].Title}
					sort.Strings(listName)

					if sortQuery != "" && sortQuery == "ASC" {
						if listName[0] == movies[i].Title {
							result = true
						} else {
							result = false
						}
					} else if sortQuery != "" && sortQuery == "DESC" {
						if listName[0] == movies[i].Title {
							result = false
						} else {
							result = true
						}
					}
				} else {
					if sortQuery != "" && sortQuery == "ASC" {
						result = (movies[i].Year < movies[j].Year)
					} else if sortQuery != "" && sortQuery == "DESC" {
						result = (movies[i].Year > movies[j].Year)
					}
				}
				return result
			})
		}
		response["result"] = movies

		fmt.Println("***********RESP************")
		fmt.Println(response)

		js, _ := json.MarshalIndent(response, "", "\t")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		//fmt.Fprintf(w, strings.ToUpper(""))
		//json.NewEncoder(w).Encode(response)
	}
}

// NewHandler returns a router with all endpoint handlers
func NewHandler(s movies.MovieSearcher) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/movies", createSearchMoviesHandler(s)).Methods("GET")
	router.HandleFunc("/movies-sorted", CreateSearchMoviesHandlerSort(s)).Methods("GET")
	return router
}
