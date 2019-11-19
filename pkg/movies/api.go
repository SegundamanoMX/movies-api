package movies

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"errors"
)

// APIMovieSearcher is a MovieSearcher implementation using omdbapi
type APIMovieSearcher struct {
	APIKey string
	URL    string
}

type omdbapiResponse struct {
	Search   []Movie `json:"Search"`
	Total    string  `json:"totalResults"`
	Response string  `json:"Response"`
	Error    string  `json:"Error"`
}

type finalRespose struct {
	Search []Movie
	total  int64
}

func (box *finalRespose) AddItem(items []Movie) []Movie {
	box.Search = append(box.Search, items...)
	return box.Search
}

func sortMovies(movies []Movie) []Movie {
	sort.SliceStable(movies, func(i, j int) bool {

		if movies[i].Year == movies[j].Year {
			return strings.ToLower(movies[i].Title) < strings.ToLower(movies[j].Title)
		} else {
			return movies[i].Year < movies[j].Year
		}
	})

	return movies

}

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMovies(query string, sort bool) ([]Movie, error) {

	var totalPages int64 = math.MaxInt64
	var page int64 = 1
	var respStruct *finalRespose
	respStruct = new(finalRespose)
	respStruct.total = 0
	respStruct.Search = []Movie{}
	var itemsPerPage int64 = 10
	var current int64 = 1
	for current >= page {
		// call omdbapi
		params := url.Values{}
		params.Add("s", query)
		params.Add("apikey", s.APIKey)
		params.Add("type", "movie")
		params.Add("page", strconv.FormatInt(page, 10))

		resp, err := http.Get(s.URL + "?" + params.Encode())
		if err != nil {
			return nil, err
		}

		// unmarshall response and get the movie array
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response omdbapiResponse
		jsonError := json.Unmarshal(respBody, &response)
		if jsonError != nil {
			return nil, jsonError
		}

		if response.Response == "False" {
			return nil, errors.New(response.Error)
		}

		respStruct.AddItem(response.Search)
		totalPages, err = strconv.ParseInt(response.Total, 10, 64)

		if page == 1 {
			d := float64(totalPages) / float64(itemsPerPage)
			current = int64(math.Ceil(d))
		}

		page++

	}

	// return result

	if sort {
		return sortMovies(respStruct.Search), nil
	} 
	
	return respStruct.Search, nil
	
}
