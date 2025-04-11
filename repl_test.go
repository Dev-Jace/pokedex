package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	// ...
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
		{
			input:    "  hello  there  ",
			expected: []string{"hello", "there"},
		},
		{
			input:    "  general kenobi   ",
			expected: []string{"general", "kenobi"},
		},
		{
			input:    "  this is a good test of strength  ",
			expected: []string{"this", "is", "a", "good", "test", "of", "strength"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}

//*/
