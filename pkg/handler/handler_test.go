package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

// const with the body to mock the request of omdbapi
const OMDBAPIMockBody = `{"Search":[{"Title":"The Fast and the Furious: Tokyo Drift","Year":"2006"},{"Title":"2 Fast 2 Furious","Year":"2003"},{"Title":"Fast Food Nation","Year":"2006"}], "totalResults": "3"}`

// Router returns a router with the endpoints to test
func Router(s movies.MovieSearcher) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/movies", createSearchMoviesHandler(s)).Methods("GET")
	router.HandleFunc("/movies-sorted", createSearchAndSortMoviesHandler(s)).Methods("GET")
	return router
}

// TestEndpoints tests the response of the endpoints of movie-api
func TestEndpoints(t *testing.T) {
	// dummy searcher
	searcher := &movies.APIMovieSearcher{
		URL:    "http://example.com/",
		APIKey: "mock-api-key",
	}
	// Mock for the request of omdbapi
	httpmock.RegisterResponder(
		"GET",
		"http://example.com/",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, OMDBAPIMockBody), nil
		},
	)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	// Cases to test
	cases := []struct {
		name                string
		uri                 string
		expectedBody        string
		expectedErrorString string
	}{
		{
			name:                "MoviesCase",
			uri:                 "/movies?q=fast",
			expectedBody:        `{"result":[{"Title":"The Fast and the Furious: Tokyo Drift","Year":"2006"},{"Title":"2 Fast 2 Furious","Year":"2003"},{"Title":"Fast Food Nation","Year":"2006"}]}`,
			expectedErrorString: "",
		},
		{
			name:                "SortedMoviesCase",
			uri:                 "/movies-sorted?q=fast",
			expectedBody:        `{"result":[{"Title":"2 Fast 2 Furious","Year":"2003"},{"Title":"Fast Food Nation","Year":"2006"},{"Title":"The Fast and the Furious: Tokyo Drift","Year":"2006"}]}`,
			expectedErrorString: "",
		},
	}
	for _, c := range cases {
		request, _ := http.NewRequest("GET", c.uri, nil)
		response := httptest.NewRecorder()
		Router(searcher).ServeHTTP(response, request)
		assert.JSONEq(t, c.expectedBody, response.Body.String(), "The movies list is sorted as expected")
	}
}
