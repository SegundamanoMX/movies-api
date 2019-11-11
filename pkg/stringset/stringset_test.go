package stringset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareStringsAlphabetically(t *testing.T) {
	type comparison struct {
		a, b     string
		expected bool
	}
	comparisons := []comparison{
		{"hi", "high", true},
		{"high", "hig", false},
		{"hi ghly", "hig", true},
		{"And", "android", true},
		{"and", "Android", true},
		{"and", "android", true},
		{"ANDroid", "and", false},
	}

	for _, c := range comparisons {
		result := CompareStringsAlphabetically(c.a, c.b)
		assert.EqualValues(t, c.expected, result, fmt.Sprintf("Comparing: '%s' With:'%s'", c.a, c.b))
	}
}
