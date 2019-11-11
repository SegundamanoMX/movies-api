//Package stringset implements functions to manipulate strings that are not included in the package strings
package stringset

import "unicode"

// CompareStringsAlphabetically compares two strings to be ordered alphabetically (ignoring case);
// returns true if a goes first, false otherwise.
func CompareStringsAlphabetically(a, b string) bool {
	ta, tb := []rune(a), []rune(b)
	aIsLonger := true
	max := len(ta)
	if lb := len(tb); lb < max {
		aIsLonger = false
		max = lb
	}

	for k := 0; k < max; k++ {
		ra, rb := unicode.ToLower(ta[k]), unicode.ToLower(tb[k])
		if ra != rb {
			return ra < rb
		}
	}
	return aIsLonger
}
