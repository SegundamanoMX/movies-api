package movies

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"

	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/stringset"
)

var (
	// ErrInvalidApikey means that the omdbapi key is invalid.
	ErrInvalidApikey = errors.New("invalid API key")
	// ErrMovieNotFound means that the query in the omdbapi didn't return any results.
	ErrMovieNotFound = errors.New("movie not found")
	// ErrToManyResults means that the omdbapi needs a more specific query (it has too many results to handle the request).
	ErrToManyResults = errors.New("too many results")
	// ErrUnknown means that the omdbapi an error not handled by this api.
	ErrUnknown = errors.New("unknown")
)

// APIMovieSearcher is a MovieSearcher implementation using omdbapi.
type APIMovieSearcher struct {
	APIKey string
	URL    string
}

type omdbapiResponse struct {
	Result           bool    `json:"Result"`
	Search           []Movie `json:"Search"`
	TotalResults     int     `json:"totalResults,string"`
	ErrorDescription string  `json:"Error"`
	Error            error   `json:"-"`
}

// SearchMovies searches for all the movies that match the given query.
func (s *APIMovieSearcher) SearchMovies(query string, sorted bool) ([]Movie, error) {
	movies, totalResults, err := s.searchMovies(query, 1)
	if err != nil {
		if err == ErrMovieNotFound {
			return nil, nil
		}
		return nil, err
	}
	pages := int(math.Ceil(float64(totalResults) / 10.0))
	
	for i := 2; i <= pages; i++ {
		moviesAux, _, err := s.searchMovies(query, i)
		if err != nil {
			return nil, err
		}
		movies = append(movies, moviesAux...)
	}
	if sorted {
		sortMovies(movies)
	}
	return movies, nil
}

// SearchMovies searches for all the movies that match the given query and page.
func (s *APIMovieSearcher) searchMovies(query string, page int) ([]Movie, int, error) {
	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
	params.Add("page", strconv.Itoa(page))
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
	err = json.Unmarshal(respBody, &respStruct)
	if err != nil {
		return nil, 0, nil
	}
	respStruct.parseErrorDescription()
	return respStruct.Search, respStruct.TotalResults, respStruct.Error
}

// sortMovies sorts the movie results by year of release (ascending), and if there is a tie, in alphabetical order by title (ascending).
func sortMovies(movies []Movie) {
	sort.Slice(movies, func(i, j int) bool {
		if movies[i].Year != movies[j].Year {
			return movies[i].Year < movies[j].Year
		}
		return stringset.CompareStringsAlphabetically(movies[i].Title, movies[j].Title)
	})
}

// parseErrorDescription parses the errorDescription into the Error field if the amdbapi returned one.
func (o *omdbapiResponse) parseErrorDescription() {
	switch o.ErrorDescription {
	case "":
		o.Error = nil
	case "Invalid API key!":
		o.Error = ErrInvalidApikey
	case "Movie not found!":
		o.Error = ErrMovieNotFound
	case "Too many results.":
		o.Error = ErrToManyResults
	default:
		o.Error = ErrUnknown
	}
}
