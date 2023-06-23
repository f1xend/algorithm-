package leetcode

import "testing"

func TestSimilarPairs(t *testing.T) {

	cases := []struct {
		Name          string
		Input         []string
		ExpectedCalls int
	}{
		{
			"first test, 5 sequence",
			[]string{"aba", "aabb", "abcd", "bac", "aabc"},
			2,
		},
		{
			"second test, 3 sequence",
			[]string{"aabb", "ab", "ba"},
			3,
		},
	}

	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			got := similarPairs(v.Input)
			want := v.ExpectedCalls

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
