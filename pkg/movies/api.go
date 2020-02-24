package movies

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// APIMovieSearcher is a MovieSearcher implementation using omdbapi
type APIMovieSearcher struct {
	APIKey string
	URL    string
}

type omdbapiResponse struct {
	Search []Movie `json:Search`
}

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMovies(query string) ([]Movie, error) {

	const maxPagesPerRequest = 100
	var sliceOfMovies []Movie
	for i := 1; i <= maxPagesPerRequest; i++ {
		// call omdbapi
		params := url.Values{}
		params.Add("s", query)
		params.Add("apikey", s.APIKey)
		params.Add("type", "movie")
		params.Add("page", strconv.Itoa(i))
		resp, err := http.Get(s.URL + "?" + params.Encode())
		if err != nil {
			return nil, err
		}

		// unmarshall response and get the movie array
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var respStruct omdbapiResponse
		json.Unmarshal(respBody, &respStruct)

		if respStruct.Search == nil {
			break
		} else {
			sliceOfMovies = append(sliceOfMovies, respStruct.Search...)
		}
	}
	// return result
	return sliceOfMovies, nil
}
