package config

import (
	"math"
	"strconv"
	"time"
)

var Count int

func Increase() {
	Count = Count + 1
}

var StartFlag int
var EndFlag int
var IsHistoryJob bool

func GetStartEndTime() (string, string) {
	if IsHistoryJob {
		return startEndModeTime()
	}
	return nowStartEndTime()
}

func startEndModeTime() (string, string) {
	baseTime := StartFlag + 3600*Count
	return strconv.Itoa(baseTime), strconv.Itoa(baseTime + 3600)
}

func Times() int {
	if IsHistoryJob {
		s := (float64(EndFlag) - float64(StartFlag)) / float64(3600)
		return int(math.Ceil(s))
	}
	return 1
}

func nowStartEndTime() (string, string) {
	baseTime := time.Now().Unix()
	return strconv.FormatInt(baseTime, 10), strconv.FormatInt(baseTime+int64(3600), 10)
}

var location, _ = time.LoadLocation("Local")

func dayTime() string {
	start, _ := GetStartEndTime()
	s, _ := strconv.ParseInt(start, 10, 64)
	tm := time.Unix(s, 0).In(location)
	return tm.Format("2006-01-02")
}
