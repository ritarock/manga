package util

import (
	"time"
)

const LAYOUT_YYYYMMDD = "20060102"

var timeNow = time.Now()

func GetBeginningOfMonth() string {
	now := timeNow
	t := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	return t.Format(LAYOUT_YYYYMMDD)
}

func GetEndOfMonth() string {
	now := timeNow
	t := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local)
	t = t.Add(-time.Minute)
	return t.Format(LAYOUT_YYYYMMDD)
}
