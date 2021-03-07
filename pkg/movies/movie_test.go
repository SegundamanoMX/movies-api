package movies

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	cases := []struct {
		name            string
		movies          Movies
		expected_movies Movies
	}{
		{
			name: "RegularSort",
			movies: Movies{
				{Title: "The Godfather: Part III", Year: "1990"},
				{Title: "The Godfather: Part II", Year: "1974"},
				{Title: "The Godfather", Year: "1972"},
			},
			expected_movies: Movies{
				{Title: "The Godfather", Year: "1972"},
				{Title: "The Godfather: Part II", Year: "1974"},
				{Title: "The Godfather: Part III", Year: "1990"},
			},
		},
		{
			name: "SameYearSort",
			movies: Movies{
				{Title: "The Matrix Revolutions", Year: "2003"},
				{Title: "The Matrix", Year: "1990"},
				{Title: "The Matrix Reloaded", Year: "2003"},
			},
			expected_movies: Movies{
				{Title: "The Matrix", Year: "1990"},
				{Title: "The Matrix Reloaded", Year: "2003"},
				{Title: "The Matrix Revolutions", Year: "2003"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Sort(Movies(c.movies))
			assert.EqualValues(t, c.movies, c.expected_movies)
		})
	}
}
