package movies

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSearchMovies(t *testing.T) {
	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []Movie
		expectedErrorString string
	}{
		{
			name:             "RegularCase",
			mockResponseBody: `{"Search":[{"Title":"Star Wars: A New Hope","Year":"1977"},{"Title":"Star Wars: The Empire Strikes Back","Year":"1980"}], "totalResults":"2", "Response":"True"}`,
			expectedMovies: []Movie{
				{Title: "Star Wars: A New Hope", Year: "1977"},
				{Title: "Star Wars: The Empire Strikes Back", Year: "1980"},
			},
			expectedErrorString: "",
		},
	}

	searcher := &APIMovieSearcher{
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

func TestSearchMoviesWithSort(t *testing.T) {
	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []Movie
		expectedErrorString string
	}{
		{
			name:             "RegularCase",
			mockResponseBody: `{"Search":[{"Title":"The Adventures of Young Indiana Jones: Attack of the Hawkmen","Year":"1995"},{"Title":"Indiana Jones and the Temple of the Forbidden Eye Ride","Year":"1995"},{"Title":"The Adventures of Young Indiana Jones: Treasure of the Peacock's Eye","Year":"1995"},{"Title": "Indiana Jones and the Raiders of the Lost Ark","Year":"1981"},{"Title": "Indiana Jones and the Kingdom of the Crystal Skull","Year":"2008"}], "totalResults":"5", "Response":"True"}`,
			expectedMovies: []Movie{
				{Title: "Indiana Jones and the Raiders of the Lost Ark", Year: "1981"},
				{Title: "Indiana Jones and the Temple of the Forbidden Eye Ride", Year: "1995"},
				{Title: "The Adventures of Young Indiana Jones: Attack of the Hawkmen", Year: "1995"},
				{Title: "The Adventures of Young Indiana Jones: Treasure of the Peacock's Eye", Year: "1995"},
				{Title: "Indiana Jones and the Kingdom of the Crystal Skull", Year: "2008"},
			},
			expectedErrorString: "",
		},
	}

	searcher := &APIMovieSearcher{
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
			actualMovies, actualError := searcher.SearchMovies("indiana jones")

			if actualMovies != nil {
				SortMoviesByYearAndTitle(actualMovies)
			}

			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})
	}
}
