package utils

import (
	"log"
	"time"
)

var Logger *log.Logger

// for logger
func Warning(args ...interface{}) {
	Logger.SetPrefix("WARNING ")
	Logger.Println(args...)
}

func Danger(args ...interface{}) {
	Logger.SetPrefix("WARNING ")
	Logger.Println(args...)
}

func GetBeginningOfMonth() string {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	return t.Format(LAYOUT_YYYYMMDD)
}

func GetEndOfMonth() string {
	now := time.Now()
	t := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local)
	t = t.Add(-time.Minute)
	return t.Format(LAYOUT_YYYYMMDD)
}
