package main

import (
	"log"
	"time"
)

var logger *log.Logger

const LAYOUT_YYYYMMDD = "20060102"

// for logger
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

func getBeginningOfMonth() string {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	return t.Format(LAYOUT_YYYYMMDD)
}

func getEndOfMonth() string {
	now := time.Now()
	t := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local)
	t.Add(-time.Minute)
	return t.Format(LAYOUT_YYYYMMDD)
}
