package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    " MoUntAin_ HeRe _ ",
			expected: []string{"mountain_", "here", "_"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLen := len(actual)
		expectedLen := len(c.expected)

		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if actualLen != expectedLen {
			t.Errorf("Actual length: %d not match expected: %d", actualLen, expectedLen)
			t.FailNow()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Actual word: %v not match expected: %v", word, expectedWord)
				t.FailNow()
			}
		}
	}
}
