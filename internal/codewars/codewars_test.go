package codewars

import "testing"

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

func TestReverseString(t *testing.T) {
	got := ReverseString("Привет!")
	want := "!тевирП"

	assertCorrectMessage(t, got, want)
}

func TestCheckForFacror(t *testing.T) {
	got := CheckForFacror(6, 3)
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGrow(t *testing.T) {
	got := Grow([]int{1, 2, 3, 4})
	want := 24

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestGreet(t *testing.T) {
	got := Greet()
	want := "hello world!"

	assertCorrectMessage(t, got, want)
}

func TestStringToNumber(t *testing.T) {
	got := StringToNumber("-7")
	want := -7

	assertCorrectNumber(t, got, want)
}
