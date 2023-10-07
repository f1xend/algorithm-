package subsequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
392. Is Subsequence
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none)
of the characters without disturbing the relative positions of the remaining characters.
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).392. Is Subsequence
*/
func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		name     string
		input_s  string
		input_t  string
		expected bool
		err      error
	}{
		{
			name:     "case 1",
			input_s:  "abc",
			input_t:  "ahbgdc",
			expected: true,
			err:      nil,
		},
		{
			name:     "case 2",
			input_s:  "axc",
			input_t:  "ahbgdc",
			expected: false,
			err:      nil,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := isSubsequence(tc.input_s, tc.input_t)
			require.Equal(t, tc.expected, got)
		})
	}
}
