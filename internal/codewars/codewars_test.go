package codewars

import (
	"reflect"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertCorrectNumber(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertCorrectBool(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseString(t *testing.T) {
	got := ReverseString("Привет!")
	want := "!тевирП"

	assertCorrectMessage(t, got, want)
}

func TestCheckForFacror(t *testing.T) {
	got := CheckForFacror(6, 3)
	want := true

	assertCorrectBool(t, got, want)
}

func TestGrow(t *testing.T) {
	got := Grow([]int{1, 2, 3, 4})
	want := 24

	assertCorrectNumber(t, got, want)
}

func TestGreeting(t *testing.T) {
	got := Greeting()
	want := "hello world!"

	assertCorrectMessage(t, got, want)
}

func TestStringToNumber(t *testing.T) {
	got := StringToNumber("-7")
	want := -7

	assertCorrectNumber(t, got, want)
}

func TestSetAlarm(t *testing.T) {
	t.Run("true ortions - employed and not on vacation ", func(t *testing.T) {
		got := SetAlarm(true, false)
		want := true

		assertCorrectBool(t, got, want)
	})

	t.Run("false ortions - employed and not on vacation ", func(t *testing.T) {
		got := SetAlarm(false, true)
		want := false

		assertCorrectBool(t, got, want)
	})
}

func TestRepeatStr(t *testing.T) {
	got := RepeatStr(6, "I")
	want := "IIIIII"

	assertCorrectMessage(t, got, want)
}

func TestDoubleSliceInt(t *testing.T) {
	got := DoubleSliceInt([]int{1, 2, 3})
	want := []int{2, 4, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// 7 kyu tasks
func TestHero(t *testing.T) {
	cases := []struct {
		Name  string
		Input struct {
			Bullets int
			Dragons int
		}
		ExpectedCalls bool
	}{
		{
			"enouth bullets to win some dragons",
			struct {
				Bullets int
				Dragons int
			}{10, 5},
			true,
		},
		{
			"not enouth bullets to win some dragons",
			struct {
				Bullets int
				Dragons int
			}{7, 4},
			false,
		},
		{
			"check 0 bullets and 1 dragon",
			struct {
				Bullets int
				Dragons int
			}{0, 1},
			false,
		},
		{
			"check 0 dragons and 1 bullet",
			struct {
				Bullets int
				Dragons int
			}{1, 0},
			true,
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			bullets := test.Input.Bullets
			dragons := test.Input.Dragons
			got := Hero(bullets, dragons)

			want := test.ExpectedCalls

			assertCorrectBool(t, got, want)
		})
	}

}
