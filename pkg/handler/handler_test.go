package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

type Respueta struct {
	Result []movies.Movie `json:"result"`
}

type Bird struct {
	Species     string
	Description string
}

func TestCreateSearchMoviesHandlerSort(t *testing.T) {

	cases := []struct {
		name                string
		expectedMovies      []movies.Movie
		url                 string
		mockResponseBody    string
		expectedErrorString string
	}{
		{
			name:             "Sort ASC",
			mockResponseBody: `{"Search":[{"Title":"The Matrix","Year":"1999"},{"Title":"The Matrix Reloaded","Year":"2003"},{"Title":"The Matrix Revolutions","Year":"2003"},{"Title":"The Matrix Resurrections","Year":"2021"},{"Title":"Making 'The Matrix'","Year":"1999"},{"Title":"The Matrix Revisited","Year":"2001"},{"Title":"A Glitch in the Matrix","Year":"2021"},{"Title":"Armitage III: Dual Matrix","Year":"2002"},{"Title":"Sex and the Matrix","Year":"2000"},{"Title":"The Matrix Reloaded Revisited","Year":"2004"}]}`,
			expectedMovies: []movies.Movie{
				{Title: "Making 'The Matrix'", Year: "1999"},
				{Title: "The Matrix", Year: "1999"},
				{Title: "Sex and the Matrix", Year: "2000"},
				{Title: "The Matrix Revisited", Year: "2001"},
				{Title: "Armitage III: Dual Matrix", Year: "2002"},
				{Title: "The Matrix Reloaded", Year: "2003"},
				{Title: "The Matrix Revolutions", Year: "2003"},
				{Title: "The Matrix Reloaded Revisited", Year: "2004"},
				{Title: "A Glitch in the Matrix", Year: "2021"},
				{Title: "The Matrix Resurrections", Year: "2021"},
			},
			expectedErrorString: "",
			url:                 "/movies-sort?q='matrix'&sort=ASC&page=1",
		},
		{
			name:             "Sort DESC",
			mockResponseBody: `{"Search":[{"Title":"The Matrix","Year":"1999"},{"Title":"The Matrix Reloaded","Year":"2003"},{"Title":"The Matrix Revolutions","Year":"2003"},{"Title":"The Matrix Resurrections","Year":"2021"},{"Title":"Making 'The Matrix'","Year":"1999"},{"Title":"The Matrix Revisited","Year":"2001"},{"Title":"A Glitch in the Matrix","Year":"2021"},{"Title":"Armitage III: Dual Matrix","Year":"2002"},{"Title":"Sex and the Matrix","Year":"2000"},{"Title":"The Matrix Reloaded Revisited","Year":"2004"}]}`,
			expectedMovies: []movies.Movie{

				{Title: "The Matrix Resurrections", Year: "2021"},
				{Title: "A Glitch in the Matrix", Year: "2021"},
				{Title: "The Matrix Reloaded Revisited", Year: "2004"},
				{Title: "The Matrix Revolutions", Year: "2003"},
				{Title: "The Matrix Reloaded", Year: "2003"},
				{Title: "Armitage III: Dual Matrix", Year: "2002"},
				{Title: "The Matrix Revisited", Year: "2001"},
				{Title: "Sex and the Matrix", Year: "2000"},
				{Title: "The Matrix", Year: "1999"},
				{Title: "Making 'The Matrix'", Year: "1999"},
			},
			expectedErrorString: "",
			url:                 "/movies-sort?q='matrix'&sort=DESC&page=1",
		},
	}

	for _, c := range cases {

		httpmock.RegisterResponder(
			"GET",
			"https://www.omdbapi.com/",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		req, errReq := http.NewRequest(http.MethodGet, c.url, nil)
		searcher := &movies.APIMovieSearcher{
			APIKey: "xxxxx",
			URL:    "https://www.omdbapi.com/",
		}

		if errReq != nil {
			t.Fatalf("Bad request")
		}

		rec := httptest.NewRecorder()
		CreateSearchMoviesHandlerSort(searcher)(rec, req)
		jsonLine, _ := rec.Body.ReadString(0)
		var payload Respueta
		json.Unmarshal([]byte(jsonLine), &payload)
		assert.EqualValues(t, c.expectedMovies, payload.Result)
	}

}
