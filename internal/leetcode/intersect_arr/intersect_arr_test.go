package intersectarr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
350. Intersection of Two Arrays II
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.
*/
func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
		err      error
	}{
		{
			name:     "case 1",
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2, 2},
			err:      nil,
		},
		{
			name:     "case 2",
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{4, 9},
			err:      nil,
		},
		{
			name:     "case 3",
			nums1:    []int{2,1},
			nums2:    []int{1,1},
			expected: []int{1},
			err:      nil,
		},
		{
			name:     "case 4",
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2},
			expected: []int{2},
			err:      nil,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := Intersect(tc.nums1, tc.nums2)
			require.Equal(t, tc.expected, got)
		})
	}
}
