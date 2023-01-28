package crawler

import (
	"reflect"
	"testing"
)

func Test_makeIsbnList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		coverage := []string{
			"1",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
			"9",
			"10",
		}
		got := makeIsbnList(coverage)
		want := [][]string{coverage}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("10000 elements in one array", func(t *testing.T) {
		coverage := make([]string, 10001)
		isbn := makeIsbnList(coverage)
		got := len(isbn[0])
		want := 10000
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
