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
	TotalResults string `json:totalResults`
}

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMovies(query string,page string) ([]Movie, int, error) {

	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
	params.Add("page", page)
	resp, err := http.Get(s.URL + "?" + params.Encode())
	if err != nil {
		return nil, 1,err
	}

	// unmarshall response and get the movie array
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,1,nil
	}
	var respStruct omdbapiResponse
	json.Unmarshal(respBody, &respStruct)
	total, err := strconv.Atoi(respStruct.TotalResults)
	// return result
	return respStruct.Search,total, nil
}
