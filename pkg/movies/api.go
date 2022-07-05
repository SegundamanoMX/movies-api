package movies

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

// APIMovieSearcher is a MovieSearcher implementation using omdbapi
type APIMovieSearcher struct {
	APIKey string
	URL    string
}

type omdbapiResponse struct {
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
	Search       []Movie `json:"Search"`
}

func (s *APIMovieSearcher) SearchMoviesByPage(query string, page int) (*omdbapiResponse, error) {

	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
	if page > 0 {
		params.Add("page", strconv.Itoa(page))
	}
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
	err = json.Unmarshal(respBody, &respStruct)
	if err != nil {
		return nil, err
	}

	// return result
	return &respStruct, nil
}

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMovies(query string) ([]Movie, error) {

	page := 0
	response, err := s.SearchMoviesByPage(query, page)
	result := response.Search
	if err != nil {
		return nil, err
	}
	total, _ := strconv.Atoi(response.TotalResults)
	for ; len(result) < total; page++ {
		response, err := s.SearchMoviesByPage(query, page)
		if err != nil {
			break
		}
		result = append(result, response.Search...)
	}

	return result, nil
}

func (s *APIMovieSearcher) SearchMoviesSorted(query string) ([]Movie, error) {

	result, err := s.SearchMovies(query)
	if err != nil {
		return nil, err
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Year == result[j].Year {
			return result[i].Title < result[j].Title
		}
		return result[i].Year < result[j].Year
	})

	return result, nil
}
