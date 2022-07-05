package movies

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/constants"
)

// APIMovieSearcher is a MovieSearcher implementation using omdbapi
type APIMovieSearcher struct {
	APIKey string
	URL    string
}

type omdbapiResponse struct {
	Search       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}

// Method that search movie by name with the provider
func getMoviesFromProvider(urlProvider string, apiKey string, movieTitleToSearch string, pageToReturn string) (omdbapiResponse, error) {
	var respStruct omdbapiResponse

	// call omdbapi
	params := url.Values{}
	params.Add(constants.PARAM_MOVIE_TITLE_TO_SEARCH, movieTitleToSearch)
	params.Add(constants.PARAM_API_KEY, apiKey)
	params.Add(constants.PARAM_TYPE, constants.MOVIE_TYPE)
	params.Add(constants.PARAM_PAGE, pageToReturn)
	resp, err := http.Get(urlProvider + constants.QUESTION_MARK + params.Encode())

	if err != nil {
		log.Println("Error movies.GetMoviesFromProvider - Get - urlProvider = [" + urlProvider +
			"], movieTitleToSearch = [" + movieTitleToSearch + "], pageToReturn = [" + pageToReturn +
			"], error = [" + err.Error() + "]")
		return respStruct, err
	}

	// unmarshall response and get the movie array
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error movies.GetMoviesFromProvider - ReadAll - urlProvider = [" + urlProvider +
			"], movieTitleToSearch = [" + movieTitleToSearch + "], pageToReturn = [" + pageToReturn +
			"], error = [" + err.Error() + "]")
		return respStruct, err
	}

	json.Unmarshal(respBody, &respStruct)

	// return result
	return respStruct, nil
}

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMovies(movieTitleToSearch string) ([]Movie, error) {
	var pageNumber int = 1
	var recoveredItems int = 0
	var totalResults int = 0
	var err error
	var movies []Movie
	var responseWS omdbapiResponse

	for {
		responseWS, err = getMoviesFromProvider(s.URL, s.APIKey, movieTitleToSearch, strconv.Itoa(pageNumber))

		if err != nil {
			return nil, err
		}

		if constants.TRUE_AS_STR == responseWS.Response {
			if totalResults == 0 {
				totalResults, err = strconv.Atoi(responseWS.TotalResults)
				if err != nil {
					log.Println("Error movies.SearchMovies - movieTitleToSearch = [" + movieTitleToSearch +
						"], error = [" + err.Error() + "]")
					return nil, err
				}
			}

			movies = append(movies, responseWS.Search[:]...)

			recoveredItems += len(responseWS.Search)
		}

		if recoveredItems == totalResults {
			break
		}

		pageNumber++
	}

	return movies, nil
}
