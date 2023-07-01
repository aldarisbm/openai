package main

import (
	"testing"
)

// TestValidateToken tests the validateToken function.

func TestValidateToken(t *testing.T) {
	// Test cases.
	testCases := []struct {
		name     string
		input    string
		expected error
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: InvalidTokenError,
		},
		{
			name:     "Invalid token",
			input:    "sk-1234567890",
			expected: InvalidTokenError,
		},
		{
			name:     "Valid token",
			input:    "sk-1234567890abcdef1234567890abcdef1234567890abcdef",
			expected: nil,
		},
	}

	// Run the test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := validateToken(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
