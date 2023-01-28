package viewer

import (
	"testing"
)

func assert(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_validateYyyy(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got := validateYyyy("2023")
		want := "2023"
		assert(t, got, want)
	})

	t.Run("under 1970", func(t *testing.T) {
		got := validateYyyy("1969")
		want := ""
		assert(t, got, want)
	})

	t.Run("failed", func(t *testing.T) {
		got := validateYyyy("a")
		want := ""
		assert(t, got, want)
	})
}

func Test_validateMm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got := validateMm("01")
		want := "01"
		assert(t, got, want)
	})

	t.Run("success 0 padding", func(t *testing.T) {
		got := validateMm("1")
		want := "01"
		assert(t, got, want)
	})

	t.Run("out of range", func(t *testing.T) {
		got := validateMm("15")
		want := ""
		assert(t, got, want)
	})

	t.Run("failed", func(t *testing.T) {
		got := validateMm("a")
		want := ""
		assert(t, got, want)
	})
}
