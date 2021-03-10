package movies

import (
	"sort"
	"strconv"
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
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

	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
	log.Print(s.URL + "?" + params.Encode())
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
	
	log.Print("Hola:")
	log.Print(respStruct.Search)

	// return result
	return respStruct.Search, nil
}

type ByYear []MovieI
func (a ByYear) Len() int           { return len(a) }
func (a ByYear) Less(i, j int) bool {
	if(a[i].Year < a[j].Year){
		return true
	}
	if(a[i].Year > a[j].Year){
		return false
	} 
	return a[i].Title < a[j].Title 
}
func (a ByYear) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMoviesSorted(query string) ([]MovieI, error) {

	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
	log.Print(s.URL + "?" + params.Encode())
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
	
	log.Print("Hola:")
	log.Print(respStruct.Search)

	peliculas := respStruct.Search

	peliculassort := []MovieI{}

	for _, pe := range peliculas {
		yeari, err := strconv.Atoi(pe.Year)
		if err != nil {
			log.Print(err.Error())
		}
		p := MovieI{pe.Title, yeari}
		peliculassort = append(peliculassort, p)
	}

	sort.Sort(ByYear(peliculassort))




	// return result
	return peliculassort, nil
}
