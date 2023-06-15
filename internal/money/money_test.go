package money

import (
	"reflect"
	"testing"
)

func TestChange(t *testing.T) {
	denomination := []int{25, 10, 5, 1}
	cases := []struct {
		Name          string
		money         Money
		ExpectedCalls []int
	}{
		{
			"for each one coin",
			16,
			[]int{10, 5, 1},
		},
		{
			"for one coin multiple times",
			75,
			[]int{25, 25, 25},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := Change(test.money, denomination)

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	// t.Run("for each one coin", func(t *testing.T) {
	// 	denomination :=

	// 	var money Money = 16

	// 	got := Change(money, denomination)
	// 	want := []int{10, 5, 1}

	// 	if !reflect.DeepEqual(got, want) {
	// 		t.Errorf("got %v, want %v", got, want)
	// 	}
	// })
}
