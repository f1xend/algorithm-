package jumpgame

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanReach(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		minJump int
		maxJump int
		expected bool
		err      error
	}{
		{
			name:     "case 1",
			input:    "011010",
			minJump: 2,
			maxJump: 3,
			expected: true,
			err:      nil,
		},
		{
			name:     "case2",
			input:    "01101110",
			minJump: 2,
			maxJump: 3,
			expected: false,
			err:      nil,
		},
		{
			name:     "case3",
			input:    "00",
			minJump: 1,
			maxJump: 1,
			expected: true,
			err:      nil,
		},
		{
			name:     "case4",
			input:    "01",
			minJump: 1,
			maxJump: 1,
			expected: false,
			err:      nil,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := canReach(tc.input, tc.minJump, tc.maxJump)
			require.Equal(t, tc.expected, got)
		})
	}
}