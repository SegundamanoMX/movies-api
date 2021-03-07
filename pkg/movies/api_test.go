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
			mockResponseBody: `{"Search":[{"Title":"Star Wars: A New Hope","Year":"1977"},{"Title":"Star Wars: The Empire Strikes Back","Year":"1980"}]}`,
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
			actualMovies, actualError := searcher.SearchMovies("star wars", "1")
			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})
	}
}

func TestSortedSearchMovies(t *testing.T) {
	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []Movie
		expectedErrorString string
	}{
		{
			name:             "RegularSortedCase1Page1",
			mockResponseBody: `{"Search": [{"Title": "Pirates of the Caribbean: Curse of the Black Pearl","Year": "2003"},{"Title": "Black Swan","Year": "2010"}]}`,
			expectedMovies: []Movie{
				{Title: "Pirates of the Caribbean: Curse of the Black Pearl", Year: "2003"},
				{Title: "Black Swan", Year: "2010"},
			},
			expectedErrorString: "",
		},
		{
			name:             "RegularSortedCase1Page2",
			mockResponseBody: `{"Search": [{"Title": "Black Mass", "Year": "2015"}, {"Title": "Black Mirror: Bandersnatch", "Year": "2018"}]}`,
			expectedMovies: []Movie{
				{Title: "Black Mass", Year: "2015"},
				{Title: "Black Mirror: Bandersnatch", Year: "2018"},
			},
			expectedErrorString: "",
		},
	}

	searcher := &APIMovieSearcher{
		URL:    "http://example.com/",
		APIKey: "mock-api-key",
	}

	for _, c := range cases {
		httpmock.RegisterResponder(
			"GET",
			"http://example.com/",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		switch {
		case c.name == "RegularSortedCase1Page1":
			t.Run(c.name, func(t *testing.T) {
				actualMovies, actualError := searcher.SearchMovies("black", "1")
				assert.EqualValues(t, c.expectedMovies, actualMovies)
				if c.expectedErrorString == "" {
					assert.NoError(t, actualError)
				} else {
					assert.EqualError(t, actualError, c.expectedErrorString)
				}
			})
		case c.name == "RegularSortedCase1Page2":
			t.Run(c.name, func(t *testing.T) {
				actualMovies, actualError := searcher.SearchMovies("black", "2")
				assert.EqualValues(t, c.expectedMovies, actualMovies)
				if c.expectedErrorString == "" {
					assert.NoError(t, actualError)
				} else {
					assert.EqualError(t, actualError, c.expectedErrorString)
				}
			})
		}
	}
}
