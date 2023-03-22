package main

import "testing"

// TestGetSummaryBetweenThreeBrackets tests the getSummaryBetweenThreeBrackets function.
func TestGetSummaryBetweenThreeBrackets(t *testing.T) {
	// Test cases.
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "No brackets",
			input:    "Hello, world!",
			expected: "",
		},
		{
			name:     "One bracket",
			input:    "Hello, [ world!]",
			expected: "",
		},
		{
			name:     "Two brackets",
			input:    "Hello, [[ world!]]",
			expected: "",
		},
		{
			name:     "Three brackets",
			input:    "Hello, [[[I'm Alive!]]]",
			expected: "I'm Alive!",
		},
	}

	// Run the test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getSummaryBetweenThreeBrackets(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
