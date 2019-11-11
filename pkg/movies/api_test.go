package movies

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestSearchMovies makes the test when the responses are correct or no movies are found.
func TestSearchMovies(t *testing.T) {
	cases := []struct {
		name              string
		sorted            bool
		mockResponsesBody []string
		expectedMovies    []Movie
	}{
		{
			name:              "RegularCase",
			sorted:            false,
			mockResponsesBody: []string{`{"Search":[{"Title":"Star Wars: A New Hope","Year":"1977"},{"Title":"Star Wars: The Empire Strikes Back","Year":"1980"}]}`},
			expectedMovies: []Movie{
				{Title: "Star Wars: A New Hope", Year: "1977"},
				{Title: "Star Wars: The Empire Strikes Back", Year: "1980"},
			},
		},
		{
			name:              "NoMoviesCase",
			sorted:            false,
			mockResponsesBody: []string{`{"Error":"Movie not found!"`},
			expectedMovies:    nil,
		},
		{
			name:   "PaginatedCase",
			sorted: false,
			mockResponsesBody: []string{
				`{"Search":[{"Title":"Star Wars: Episode IV - A New Hope","Year":"1977"},{"Title":"Star Wars: Episode V - The Empire Strikes Back","Year":"1980"},{"Title":"Star Wars: Episode VI - Return of the Jedi","Year":"1983"},{"Title":"Star Wars: Episode VII - The Force Awakens","Year":"2015"},{"Title":"Star Wars: Episode I - The Phantom Menace","Year":"1999"},{"Title":"Star Wars: Episode III - Revenge of the Sith","Year":"2005"},{"Title":"Star Wars: Episode II - Attack of the Clones","Year":"2002"},{"Title":"Rogue One: A Star Wars Story","Year":"2016"},{"Title":"Star Wars: Episode VIII - The Last Jedi","Year":"2017"},{"Title":"Solo: A Star Wars Story","Year":"2018"}],"totalResults":"21"}`,
				`{"Search":[{"Title":"Star Wars: The Clone Wars","Year":"2008"},{"Title":"The Star Wars Holiday Special","Year":"1978"},{"Title":"Robot Chicken: Star Wars","Year":"2007"},{"Title":"Robot Chicken: Star Wars Episode II","Year":"2008"},{"Title":"Robot Chicken: Star Wars III","Year":"2010"},{"Title":"Empire of Dreams: The Story of the 'Star Wars' Trilogy","Year":"2004"},{"Title":"Lego Star Wars: The Padawan Menace","Year":"2011"},{"Title":"Star Wars: Revelations","Year":"2005"},{"Title":"Lego Star Wars: The Empire Strikes Out","Year":"2012"},{"Title":"Rogue One: A Star Wars Story - World Premiere","Year":"2016"}],"totalResults":"21"}`,
				`{"Search":[{"Title":"Star Wars: The Legacy Revealed","Year":"2007"}],"totalResults":"21"}`,
			},
			expectedMovies: []Movie{
				{Title: "Star Wars: Episode IV - A New Hope", Year: "1977"}, {Title: "Star Wars: Episode V - The Empire Strikes Back", Year: "1980"}, {Title: "Star Wars: Episode VI - Return of the Jedi", Year: "1983"}, {Title: "Star Wars: Episode VII - The Force Awakens", Year: "2015"}, {Title: "Star Wars: Episode I - The Phantom Menace", Year: "1999"}, {Title: "Star Wars: Episode III - Revenge of the Sith", Year: "2005"}, {Title: "Star Wars: Episode II - Attack of the Clones", Year: "2002"}, {Title: "Rogue One: A Star Wars Story", Year: "2016"}, {Title: "Star Wars: Episode VIII - The Last Jedi", Year: "2017"}, {Title: "Solo: A Star Wars Story", Year: "2018"}, {Title: "Star Wars: The Clone Wars", Year: "2008"}, {Title: "The Star Wars Holiday Special", Year: "1978"}, {Title: "Robot Chicken: Star Wars", Year: "2007"}, {Title: "Robot Chicken: Star Wars Episode II", Year: "2008"}, {Title: "Robot Chicken: Star Wars III", Year: "2010"}, {Title: "Empire of Dreams: The Story of the 'Star Wars' Trilogy", Year: "2004"}, {Title: "Lego Star Wars: The Padawan Menace", Year: "2011"}, {Title: "Star Wars: Revelations", Year: "2005"}, {Title: "Lego Star Wars: The Empire Strikes Out", Year: "2012"}, {Title: "Rogue One: A Star Wars Story - World Premiere", Year: "2016"}, {Title: "Star Wars: The Legacy Revealed", Year: "2007"},
			},
		},
		{
			name:   "PaginatedSortedCase",
			sorted: true,
			mockResponsesBody: []string{
				`{"Search":[{"Title":"Star Wars: Episode IV - A New Hope","Year":"1977"},{"Title":"Star Wars: Episode V - The Empire Strikes Back","Year":"1980"},{"Title":"Star Wars: Episode VI - Return of the Jedi","Year":"1983"},{"Title":"Star Wars: Episode VII - The Force Awakens","Year":"2015"},{"Title":"Star Wars: Episode I - The Phantom Menace","Year":"1999"},{"Title":"Star Wars: Episode III - Revenge of the Sith","Year":"2005"},{"Title":"Star Wars: Episode II - Attack of the Clones","Year":"2002"},{"Title":"Rogue One: A Star Wars Story","Year":"2016"},{"Title":"Star Wars: Episode VIII - The Last Jedi","Year":"2017"},{"Title":"Solo: A Star Wars Story","Year":"2018"}],"totalResults":"21"}`,
				`{"Search":[{"Title":"Star Wars: The Clone Wars","Year":"2008"},{"Title":"The Star Wars Holiday Special","Year":"1978"},{"Title":"Robot Chicken: Star Wars","Year":"2007"},{"Title":"Robot Chicken: Star Wars Episode II","Year":"2008"},{"Title":"Robot Chicken: Star Wars III","Year":"2010"},{"Title":"Empire of Dreams: The Story of the 'Star Wars' Trilogy","Year":"2004"},{"Title":"Lego Star Wars: The Padawan Menace","Year":"2011"},{"Title":"Star Wars: Revelations","Year":"2005"},{"Title":"Lego Star Wars: The Empire Strikes Out","Year":"2012"},{"Title":"Rogue One: A Star Wars Story - World Premiere","Year":"2016"}],"totalResults":"21"}`,
				`{"Search":[{"Title":"Star Wars: The Legacy Revealed","Year":"2007"}],"totalResults":"21"}`,
			},
			expectedMovies: []Movie{{Title: "Star Wars: Episode IV - A New Hope", Year: "1977"}, {Title: "The Star Wars Holiday Special", Year: "1978"}, {Title: "Star Wars: Episode V - The Empire Strikes Back", Year: "1980"}, {Title: "Star Wars: Episode VI - Return of the Jedi", Year: "1983"}, {Title: "Star Wars: Episode I - The Phantom Menace", Year: "1999"}, {Title: "Star Wars: Episode II - Attack of the Clones", Year: "2002"}, {Title: "Empire of Dreams: The Story of the 'Star Wars' Trilogy", Year: "2004"}, {Title: "Star Wars: Episode III - Revenge of the Sith", Year: "2005"}, {Title: "Star Wars: Revelations", Year: "2005"}, {Title: "Robot Chicken: Star Wars", Year: "2007"}, {Title: "Star Wars: The Legacy Revealed", Year: "2007"}, {Title: "Robot Chicken: Star Wars Episode II", Year: "2008"}, {Title: "Star Wars: The Clone Wars", Year: "2008"}, {Title: "Robot Chicken: Star Wars III", Year: "2010"}, {Title: "Lego Star Wars: The Padawan Menace", Year: "2011"}, {Title: "Lego Star Wars: The Empire Strikes Out", Year: "2012"}, {Title: "Star Wars: Episode VII - The Force Awakens", Year: "2015"}, {Title: "Rogue One: A Star Wars Story", Year: "2016"}, {Title: "Rogue One: A Star Wars Story - World Premiere", Year: "2016"}, {Title: "Star Wars: Episode VIII - The Last Jedi", Year: "2017"}, {Title: "Solo: A Star Wars Story", Year: "2018"}},
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
				// get parameter from request
				queryParams := req.URL.Query()
				page, _ := strconv.Atoi(queryParams.Get("page"))
				return httpmock.NewStringResponse(200, c.mockResponsesBody[page-1]), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.SearchMovies("star wars", c.sorted)
			assert.EqualValues(t, c.expectedMovies, actualMovies)
			assert.NoError(t, actualError)
		})
	}
}

// TestErrorResponses tests cases where posible error may occur.
func TestErrorResponses(t *testing.T) {
	cases := []struct {
		name                string
		sorted              bool
		mockResponsesBody   []string
		mockResponsesStatus []int
		expectedMovies      []Movie
		expectedErrorString string
	}{
		{
			name:                "InvalidApiKeyCase",
			sorted:              false,
			mockResponsesBody:   []string{`{"Error":"Invalid API key!"}`},
			mockResponsesStatus: []int{401},
			expectedMovies:      nil,
			expectedErrorString: ErrInvalidApikey.Error(),
		},
		{
			name:                "TooManyResultsCase",
			sorted:              false,
			mockResponsesBody:   []string{`{"Error":"Too many results."}`},
			mockResponsesStatus: []int{200},
			expectedMovies:      nil,
			expectedErrorString: ErrToManyResults.Error(),
		},
		{
			name:                "UnknownErrCase",
			sorted:              false,
			mockResponsesBody:   []string{`{"Error":"ombdapi breaking change"}`},
			mockResponsesStatus: []int{200},
			expectedMovies:      nil,
			expectedErrorString: ErrUnknown.Error(),
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
				// get parameter from request
				return httpmock.NewStringResponse(c.mockResponsesStatus[0], c.mockResponsesBody[0]), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.SearchMovies("star wars", c.sorted)
			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})
	}
}
