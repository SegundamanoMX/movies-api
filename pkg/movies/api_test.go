package movies

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSearchMovies(t *testing.T) {
	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []Movie
		expectedErrorString string
	}{
		{
			name:             "RegularCase",
			mockResponseBody: `{"Search":[{"Title":"Star Wars: A New Hope","Year":"1977"},{"Title":"Star Wars: The Empire Strikes Back","Year":"1980"}]}`,
			expectedMovies: []Movie{
				{Title: "Star Wars: A New Hope", Year: "1977"},
				{Title: "Star Wars: The Empire Strikes Back", Year: "1980"},
			},
			expectedErrorString: "",
		},
	}
	
	searcher := &APIMovieSearcher{
		URL:    "http://example.com/",
		APIKey: "mock-api-key",
	}

	for _, c := range cases {
		// register http mock
		httpmock.RegisterResponder(
			"GET",
			"http://example.com/",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.SearchMovies("star wars") 		
			assert.EqualValues(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})
	}
}


func TestOrderMovies(t *testing.T){
	cases := []struct {
		name                string
		mockResponseBody    string
		expectedMovies      []Movie
		expectedErrorString string
	}{
		{
			name:             "OrdererCase LOFTR",
			mockResponseBody: `{"Search":[
				{"Title":"The Lord of the Rings: The Fellowship of the Ring","Year":"2001"},
				{"Title":"The Lord of the Rings: The Return of the King","Year":"2003"},
				{"Title":"The Lord of the Rings: The Two Towers","Year":"2002"},
				{"Title":"Lord of the Rings","Year":"1978"},
				{"Title":"The Making of 'The Lord of the Rings'","Year":"2002"},
				{"Title":"A Passage to Middle-earth: The Making of 'Lord of the Rings'","Year":"2001"},
				{"Title":"National Geographic: Beyond the Movie - The Lord of the Rings: Return of the King","Year":"2003"},
				{"Title":"Master of the Rings: The Unauthorized Story Behind J.R.R. Tolkien's 'Lord of the Rings'","Year":"2001"},
				{"Title":"Creating the Lord of the Rings Symphony: A Composer's Journey Through Middle-Earth","Year":"2004"},
				{"Title":"The Lord of the Rings: The Quest Fulfilled","Year":"2003"}
			]}`,
			expectedMovies: []Movie{
				{Title:"Lord of the Rings",Year:"1978"},
				{Title:"A Passage to Middle-earth: The Making of 'Lord of the Rings'",Year:"2001"},
				{Title:"Master of the Rings: The Unauthorized Story Behind J.R.R. Tolkien's 'Lord of the Rings'",Year:"2001"},
				{Title:"The Lord of the Rings: The Fellowship of the Ring",Year:"2001"},
				{Title:"The Lord of the Rings: The Two Towers",Year:"2002"},
				{Title:"The Making of 'The Lord of the Rings'",Year:"2002"},
				{Title:"National Geographic: Beyond the Movie - The Lord of the Rings: Return of the King",Year:"2003"},
				{Title:"The Lord of the Rings: The Quest Fulfilled",Year:"2003"},
				{Title:"The Lord of the Rings: The Return of the King",Year:"2003"},
				{Title:"Creating the Lord of the Rings Symphony: A Composer's Journey Through Middle-Earth",Year:"2004"},			
			},
			expectedErrorString: "",
		},

		{
			name:             "OrdererCase God",
			mockResponseBody: `{"Search":[{"Title":"City of God","Year":"2002"},
			{"Title":"Only God Forgives","Year":"2013"},
			{"Title":"God Bless America","Year":"2011"},
			{"Title":"Aguirre, the Wrath of God","Year":"1972"},
			{"Title":"OMG: Oh My God!","Year":"2012"},
			{"Title":"Son of God","Year":"2014"},
			{"Title":"Armour of God","Year":"1986"},
			{"Title":"White God","Year":"2014"},
			{"Title":"Children of a Lesser God","Year":"1986"},
			{"Title":"Astérix and Obélix: God Save Britannia","Year":"2012"}]}`,
			expectedMovies: []Movie{
				{Title:"Aguirre, the Wrath of God",Year:"1972"},
				{Title:"Armour of God",Year:"1986"},
				{Title:"Children of a Lesser God",Year:"1986"},
				{Title:"City of God",Year:"2002"},
				{Title:"God Bless America",Year:"2011"},
				{Title:"Astérix and Obélix: God Save Britannia",Year:"2012"},
				{Title:"OMG: Oh My God!",Year:"2012"},
				{Title:"Only God Forgives",Year:"2013"},
				{Title:"Son of God",Year:"2014"},
				{Title:"White God",Year:"2014"},		
			},		
			expectedErrorString: "",
		},

	}
	searcher := &APIMovieSearcher{
		URL:    "http://example.com/",
		APIKey: "mock-api-key",
	}

	for _, c := range cases {
		// register http mock
		httpmock.RegisterResponder(
			"GET",
			"http://example.com/",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, c.mockResponseBody), nil
			},
		)
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// run test
		t.Run(c.name, func(t *testing.T) {
			actualMovies, actualError := searcher.OrdererMovies("")
			assert.Equal(t, c.expectedMovies, actualMovies)
			if c.expectedErrorString == "" {
				assert.NoError(t, actualError)
			} else {
				assert.EqualError(t, actualError, c.expectedErrorString)
			}
		})
	}

}