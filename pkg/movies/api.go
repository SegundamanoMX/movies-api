package movies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// APIMovieSearcher is a MovieSearcher implementation using omdbapi
type APIMovieSearcher struct {
	APIKey string
	URL    string
}

// omdbapiResponse is the struct to get the response from omdbapi
type omdbapiResponse struct {
	Search       []Movie
	TotalResults string
}

// tractableResponse response created according to the result of the response from omdbapi
type tractableResponse struct {
	bytes  []byte
	movies []Movie
	err    error
}

const FirstPage = 1
const PageSize = 10

// SearchMovies searches for the movies using channels for the paginated requests to omdbapi
func (s *APIMovieSearcher) SearchMovies(query string) ([]Movie, error) {
	ch := make(chan tractableResponse)
	// Get a firs response to know the totals of results
	firstResponse := s.requestMoviesByQueryAndPage(query, FirstPage)
	if firstResponse.err != nil {
		return firstResponse.movies, firstResponse.err
	}
	var respStruct omdbapiResponse
	jsonErr := json.Unmarshal(firstResponse.bytes, &respStruct)
	if jsonErr != nil {
		return firstResponse.movies,
		fmt.Errorf("error with the body response of the first page: %w", jsonErr)
	}

	totalResults, totalErr := strconv.Atoi(respStruct.TotalResults)
	if totalErr != nil {
		// Default value to warranty to deliver at least the first request result
		totalResults = 0
	}
	respStruct.Search = append(respStruct.Search, s.processSubsequentPages(query, totalResults, ch)...)

	// return result
	return respStruct.Search, nil
}

// processSubsequentPages processes subsequent pages to the first
func (s *APIMovieSearcher) processSubsequentPages(query string, totalResults int, ch chan tractableResponse) []Movie {
	var pageMovies []Movie
	// Make next requests using goroutines to make them concurrently
	for i := FirstPage + 1; (i * PageSize) < totalResults; i++ {
		go s.processByChannel(query, i, ch)
		responseByChanel := <-ch
		if responseByChanel.err != nil {
			log.Printf("Error with the response of the page %d :", i)
			continue
		}
		var respNStruct omdbapiResponse
		jsonErr := json.Unmarshal(responseByChanel.bytes, &respNStruct)
		if jsonErr != nil {
			log.Fatalf("Error with the body response of the page %d :", i)
		} else {
			pageMovies = append(pageMovies, respNStruct.Search...)
		}
	}
	return pageMovies
}

// processByChannel calls the method requestMoviesByQueryAndPage using a channel
func (s *APIMovieSearcher) processByChannel(query string, page int, ch chan tractableResponse) {
	ch <- s.requestMoviesByQueryAndPage(query, page)
}

// requestMoviesByQueryAndPage make a request to omdbapi by query and page
func (s *APIMovieSearcher) requestMoviesByQueryAndPage(query string, page int) tractableResponse {
	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
	params.Add("page", strconv.Itoa(page))

	resp, err := http.Get(s.URL + "?" + params.Encode())
	if err != nil {
		return tractableResponse{nil, nil, err}
	}

	// unmarshall response and get the movie array
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tractableResponse{nil, nil, err}
	}
	return tractableResponse{respBody, nil, nil}
}
