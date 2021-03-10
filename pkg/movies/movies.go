package movies

// Movie represents a single movie
type Movie struct {
	Title string `json:"Title"`
	Year  string `json:"Year"`
}

type MovieI struct {
	Title string 
	Year  int
}

// MovieSearcher is the interfaces for anything that searches for movies
type MovieSearcher interface {
	SearchMovies(query string) ([]Movie, error)
	SearchMoviesSorted(query string) ([]MovieI, error)
}

// MovieSearcher is the interfaces for anything that searches for movies
//type MovieSearcherSorted interface {
//	SearchMoviesSorted(query string) ([]MovieI, error)
//}
