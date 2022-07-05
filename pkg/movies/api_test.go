package movies_test

import (
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSearchMovies(t *testing.T) {
	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []movies.Movie
		expectedErrorString string
	}{
		{
			name:             "RegularCase",
			mockResponseBody: `{"Search":[{"Title":"Star Wars: A New Hope","Year":"1977"},{"Title":"Star Wars: The Empire Strikes Back","Year":"1980"}]}`,
			expectedMovies: []movies.Movie{
				{Title: "Star Wars: A New Hope", Year: "1977"},
				{Title: "Star Wars: The Empire Strikes Back", Year: "1980"},
			},
			expectedErrorString: "",
		},
	}

	searcher := &movies.APIMovieSearcher{
		URL:    "http://example.com/",
		APIKey: "mock-api-key",
	}

	for _, c := range cases {
		// register http mock
		httpmock.RegisterResponder(
			"GET",
			"http://example.com/",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.SearchMovies("star wars")
			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})
	}
}

func TestSearchMoviesSorted(t *testing.T) {
	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []movies.Movie
		expectedErrorString string
	}{
		{
			name:             "Search sorted",
			mockResponseBody: `{"Search":[{"Title":"Plastic Galaxy: The Story of Star Wars Toys","Year":"2014","imdbID":"tt3648510","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMjIwNTE0MzE2Nl5BMl5BanBnXkFtZTgwNDQ1NTc0MjE@._V1_SX300.jpg"},{"Title":"The Secret Story of Toys","Year":"2014","imdbID":"tt7449658","Type":"movie","Poster":"N/A"},{"Title":"Story of the Christmas Toys as told by Mel Torme","Year":"1990","imdbID":"tt11309078","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMDk5ZjY5MmQtMTY4ZS00YThlLTkwYmItZTMxYThlODgzYTY3XkEyXkFqcGdeQXVyNTM3NDk5MTg@._V1_SX300.jpg"}],"totalResults":"3","Response":"True"}`,
			expectedMovies: []movies.Movie{
				{Title: "Story of the Christmas Toys as told by Mel Torme", Year: "1990"},
				{Title: "Plastic Galaxy: The Story of Star Wars Toys", Year: "2014"},
				{Title: "The Secret Story of Toys", Year: "2014"},
			},
			expectedErrorString: "",
		},
	}

	searcher := &movies.APIMovieSearcher{
		URL:    "http://example.com/",
		APIKey: "mock-api-key",
	}

	for _, c := range cases {
		// register http mock
		httpmock.RegisterResponder(
			"GET",
			"http://example.com/",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.SearchMoviesSorted("Toys story")
			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})
	}
}
