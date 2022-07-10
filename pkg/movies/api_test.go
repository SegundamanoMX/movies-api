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
		page                string
		search              string
		expectedErrorString string
	}{
		{
			name:             "RegularCase",
			mockResponseBody: `{"Search":[{"Title":"Star Wars: A New Hope","Year":"1977"},{"Title":"Star Wars: The Empire Strikes Back","Year":"1980"}]}`,
			expectedMovies: []Movie{
				{Title: "Star Wars: A New Hope", Year: "1977"},
				{Title: "Star Wars: The Empire Strikes Back", Year: "1980"},
			},
			page:                "1",
			search:              "star wars",
			expectedErrorString: "",
		},
		{
			name:                "No query value",
			mockResponseBody:    `{"Search":[]}`,
			expectedMovies:      []Movie{},
			page:                "1",
			search:              "",
			expectedErrorString: "",
		},
		{
			name:                "Search not found",
			mockResponseBody:    `{"Search":[]}`,
			expectedMovies:      []Movie{},
			page:                "1",
			search:              "****",
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
			actualMovies, actualError := searcher.SearchMovies(c.search, c.page)

			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})
	}
}

func TestSearchMoviesSort(t *testing.T) {

	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []Movie
		page                string
		search              string
		expectedErrorString string
	}{
		{
			name:             "Sort ASC",
			mockResponseBody: `{"Search":[{"Title":"The Matrix","Year":"1999"},{"Title":"The Matrix Reloaded","Year":"2003"},{"Title":"The Matrix Revolutions","Year":"2003"},{"Title":"The Matrix Resurrections","Year":"2021"},{"Title":"Making 'The Matrix'","Year":"1999"},{"Title":"The Matrix Revisited","Year":"2001"},{"Title":"A Glitch in the Matrix","Year":"2021"},{"Title":"Armitage III: Dual Matrix","Year":"2002"},{"Title":"Sex and the Matrix","Year":"2000"},{"Title":"The Matrix Reloaded Revisited","Year":"2004"}]}`,
			expectedMovies: []Movie{
				{Title: "The Godfather", Year: "1972"},
				{Title: "The Silence of the Lambs", Year: "1991"},
			},
			page:                "1",
			search:              "The",
			expectedErrorString: "",
		},
		{
			name:             "Sort DESC",
			mockResponseBody: `{"Search":[{"Title":"The Dark Knight Rises","Year":"2012"},{"Title":"The Avengers","Year":"2012"}]}`,
			expectedMovies: []Movie{
				{Title: "The Dark Knight Rises", Year: "2012"},
				{Title: "The Avengers", Year: "2012"},
			},
			page:                "1",
			search:              "The",
			expectedErrorString: "",
		},
		{
			name:             "SORT ASC PAGE 2",
			mockResponseBody: `{"Search":[{"Title":"The Godfather Part II","Year":"1974"},{"Title":"Star Wars: Episode V - The Empire Strikes Back","Year":"1980"}]}`,
			expectedMovies: []Movie{
				{Title: "The Godfather Part II", Year: "1974"},
				{Title: "Star Wars: Episode V - The Empire Strikes Back", Year: "1980"},
			},
			page:                "2",
			search:              "The",
			expectedErrorString: "",
		},
	}

	searcher := &APIMovieSearcher{
		URL:    "http://localhost:3000/movies-sorted/",
		APIKey: "mock-api-key",
	}

	for _, c := range cases {
		httpmock.RegisterResponder(
			"GET",
			"http://localhost:3000/movies-sorted/",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.SearchMovies(c.search, c.page)
			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})

	}

}

func TestSearchMoviesBadUrl(t *testing.T) {

	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []Movie
		page                string
		search              string
		expectedErrorString string
	}{
		{
			name:                "Bad Url",
			mockResponseBody:    `{"Search":[]}`,
			expectedMovies:      nil,
			page:                "1",
			search:              "The",
			expectedErrorString: "Get 'http://localhost:3000/movies-sorted/?apikey=mock-api-key&page=1&s=The&type=movie': no responder found",
		},
	}

	searcher := &APIMovieSearcher{
		URL:    "http://localhost:3000/movies-sorted/",
		APIKey: "mock-api-key",
	}

	for _, c := range cases {
		httpmock.RegisterResponder(
			"GET",
			"http://localhost:3000/movies-sorte",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.SearchMovies(c.search, c.page)
			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})

	}

}
