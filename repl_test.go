package main

import (
	"testing"
)

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
			input:    "bing",
			expected: []string{"bing"},
		},
		{
			input:    "base case",
			expected: []string{"base", "case"},
		},
		{
			input:    "HELLO WORLD PIKACHU",
			expected: []string{"hello", "world", "pikachu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("split slices do not match")
		}
		for i := range actual {
			word := actual[i]
			if word != c.expected[i] {
				t.Errorf("Actual word %s does not match with expected word %s", word, c.expected[i])
			}
		}
	}
}
