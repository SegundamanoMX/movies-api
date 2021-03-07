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
	Search       Movies `json:Search`
	TotalResults string `json:totalResults`
}

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMovies(query, page string) (Movies, int, error) {

	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
	params.Add("page", page)
	resp, err := http.Get(s.URL + "?" + params.Encode())
	if err != nil {
		return nil, 0, err
	}

	// unmarshall response and get the movie array
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	var respStruct omdbapiResponse
	json.Unmarshal(respBody, &respStruct)
	totalResults, _ := strconv.Atoi(respStruct.TotalResults)

	// return result
	return respStruct.Search, totalResults, nil
}
