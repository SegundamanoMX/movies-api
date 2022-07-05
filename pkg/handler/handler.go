package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/constants"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

func FetchMovies(s movies.MovieSearcher, sortResponse bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set(constants.CONTENT_TYPE, constants.APPLICATION_JSON_CONTENT_TYPE)

		var moviesArray []movies.Movie
		var err error

		// get parameter from request
		movieTitleToSearch := req.URL.Query().Get(constants.WS_PARAM_MOVIE_TITLE_TO_SEARCH)

		if movieTitleToSearch == constants.EMPTY {
			err = errors.New(constants.WS_PARAM_MOVIE_TITLE_TO_SEARCH_ERROR)
		} else {
			// call service to get movies
			moviesArray, err = s.SearchMovies(movieTitleToSearch)
		}

		// generate response
		response := make(map[string]interface{})

		if err != nil {
			response[constants.ERROR] = err.Error()
		} else {
			if moviesArray != nil && sortResponse {
				movies.SortMoviesByYearAndTitle(moviesArray)
			}
			response[constants.RESULTS] = moviesArray
		}

		json.NewEncoder(w).Encode(response)
	}
}

// ConfigRouter returns a router with all endpoint handlers
func ConfigRouter(s movies.MovieSearcher) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc(constants.ROUTE_MOVIES, FetchMovies(s, false)).Methods(http.MethodGet)
	router.HandleFunc(constants.ROUTE_MOVIES_SORTED, FetchMovies(s, true)).Methods(http.MethodGet)
	return router
}
