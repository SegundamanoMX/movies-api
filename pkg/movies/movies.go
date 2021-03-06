package movies

import "strconv"

// Movie represents a single movie
type Movie struct {
	Title string `json:"Title"`
	Year  string `json:"Year"`
}

// Movies represent a collection of Movie
type Movies []Movie

// Sort Interface implementation
func (m Movies) Len() int {
	return len(m)
}

func (m Movies) Less(i, j int) bool {
	m_i, _ := strconv.Atoi(m[i].Year)
	m_j, _ := strconv.Atoi(m[j].Year)
	if m_i < m_j {
		return true
	} else if m_i > m_j {
		return false
	}
	return m[i].Title < m[j].Title
}

func (m Movies) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// MovieSearcher is the interfaces for anything that searches for movies
type MovieSearcher interface {
	SearchMovies(query, page string) (Movies, error)
}
