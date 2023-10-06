package firstletter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
2351. First Letter to Appear Twice

Given a string s consisting of lowercase English letters, return the first letter to appear twice.

Note:
A letter a appears twice before another letter b if the second occurrence of a is before the second occurrence of b.
s will contain at least one letter that appears twice.
*/
func TestRepeatedCharacter(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected byte
		err      error
	}{
		{
			name:     "case 1",
			input:    "abccbaacz",
			expected: 'c',
			err:      nil,
		},
		{
			name:     "case2",
			input:    "abcdd",
			expected: 'd',
			err:      nil,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := repeatedCharacter(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
