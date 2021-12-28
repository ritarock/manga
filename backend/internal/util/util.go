package util

import (
	"strconv"
	"strings"
	"time"
)

const LAYOUT_YYYYMMDD = "20060102"

func GetBeginningOfMonth(date string) string {
	yyyy, mm := formatDate(date)
	t := time.Date(yyyy, time.Month(mm), 1, 0, 0, 0, 0, time.Local)
	return t.Format(LAYOUT_YYYYMMDD)
}

func GetEndOfMonth(date string) string {
	yyyy, mm := formatDate(date)
	t := time.Date(yyyy, time.Month(mm)+1, 1, 0, 0, 0, 0, time.Local)
	t = t.Add(-time.Minute)
	return t.Format(LAYOUT_YYYYMMDD)
}

func formatDate(date string) (int, int) {
	yyyy, _ := strconv.Atoi(strings.Join(strings.Split(date, "")[0:4], ""))
	mm, _ := strconv.Atoi(strings.Join(strings.Split(date, "")[4:6], ""))
	return yyyy, mm
}
