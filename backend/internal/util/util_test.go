package util

import (
	"testing"
	"time"
)

func TestGetBeginningOfMonth(t *testing.T) {
	timeNow = time.Date(2021, 1, 10, 0, 0, 0, 0, time.Local)
	got := GetBeginningOfMonth()
	want := "20210101"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetEndOfMonth(t *testing.T) {
	timeNow = time.Date(2021, 1, 10, 0, 0, 0, 0, time.Local)
	got := GetEndOfMonth()
	want := "20210131"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
