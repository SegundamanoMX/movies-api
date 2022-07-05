package movies

import (
	"sort"
)

//Method to sort movies array by Year and Title in ascending order
func SortMoviesByYearAndTitle(movies []Movie) {
	sort.Slice(movies, func(i, j int) bool {
		var sortedByYear bool = movies[i].Year < movies[j].Year

		if movies[i].Year == movies[j].Year {
			return movies[i].Title < movies[j].Title
		}

		return sortedByYear
	})
}
