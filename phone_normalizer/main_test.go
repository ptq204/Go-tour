package main

import (
	"testing"
)

func TestNormalizer(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
	}

	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			actual := normalize(c.input)
			if actual != c.want {
				t.Errorf("got %s but want %s", actual, c.want)
			}
		})
	}
}
