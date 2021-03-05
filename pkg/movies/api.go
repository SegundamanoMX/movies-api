package movies

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
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

	// return result
	return respStruct.Search, nil
}
//OrdererMovies is the method for ordering movies
func (s *APIMovieSearcher) OrdererMovies(query string) ([]Movie, error){
	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
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

	orderData :=make(map[int]map[string]string)

		for i := range respStruct.Search{		
			year , _ := strconv.Atoi(respStruct.Search[i].Year)
			_, ok := orderData[year] //check if we need a new bucket 
			if ok == false{   
				orderData[year] = map[string]string{} //new bucket			
				orderData[year][strings.ToLower(respStruct.Search[i].Title)] = respStruct.Search[i].Title								
			}else{
				orderData[year][strings.ToLower(respStruct.Search[i].Title)] = respStruct.Search[i].Title							
			}						
		}

		//format data 
		movieData := make([]Movie,0)

		//filling data into response
		keys := make([]int, 0, len(orderData))
		for k := range orderData {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		
		for _, k := range keys {
			movieKeys := make([]string, 0, len(orderData[k]))
			for movie := range orderData[k]{
				movieKeys= append(movieKeys, movie)
			}
			sort.Strings(movieKeys)
			for _, movie := range movieKeys{
				movie:= Movie{Title: orderData[k][movie], Year: strconv.Itoa(k)	}
				movieData = append(movieData, movie)
			}
		}
		json.Marshal(movieData)
		return movieData,nil		
}

