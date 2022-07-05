package movies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// APIMovieSearcher is a MovieSearcher implementation using omdbapi
type APIMovieSearcher struct {
	APIKey string
	URL    string
}

type omdbapiResponse struct {
	Search       []Movie `json:Search`
	TotalResults string  `json:totalResults`
}

type omdbapiResponseCount struct {
	TotalResults string `json:totalResults`
}

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMovies(query string, page int) ([]Movie, error) {

	var pageRequest int
	if page <= 0 {
		page = 1
	} else {
		pageRequest = page
	}
	// call omdbapi
	var params = fmt.Sprintf("s=%s&apikey=%s&type=%s&page=%s", query, s.APIKey, "movie", strconv.Itoa(pageRequest))
	resp, err := http.Get(s.URL + "?" + params)

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

	// return result
	return respStruct.Search, nil
}

// count all movies
func (s *APIMovieSearcher) SearchCount(query string) (int, error) {

	// call omdbapi

	var params = fmt.Sprintf("s=%s&apikey=%s&type=%s", query, s.APIKey, "movie")
	resp, err := http.Get(s.URL + "?" + params)
	if err != nil {
		return 0, err
	}

	// unmarshall response and get the movie array
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var respStruct omdbapiResponseCount
	json.Unmarshal(respBody, &respStruct)

	// return result
	intCount, _ := strconv.Atoi(respStruct.TotalResults)
	return intCount, nil
}
