package util

import (
	"testing"
)

func TestGetBeginningOfMonth(t *testing.T) {
	got := GetBeginningOfMonth("202112")
	want := "20211201"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetEndOfMonth(t *testing.T) {
	got := GetEndOfMonth("202112")
	want := "20211231"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_formatDate(t *testing.T) {
	gotYear, gotMonth := formatDate("202112")
	wantYear, wantMonth := 2021, 12
	if gotYear != wantYear {
		t.Errorf("gotYear %v, wantYear %v", gotYear, wantYear)
	}
	if gotMonth != wantMonth {
		t.Errorf("gotMonth %v, wantMonth %v", gotMonth, wantMonth)
	}
}
