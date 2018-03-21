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
var location, _ = time.LoadLocation("Local")

func GetStartEndTime() (string, string) {
	if IsHistoryJob {
		return startEndModeTime()
	}
	return nowStartEndTime()
}

func startEndModeTime() (string, string) {
	baseTime := StartFlag + UnixStep*Count
	return strconv.Itoa(baseTime), strconv.Itoa(baseTime + UnixStep)
}

func Cycles() int {
	if IsHistoryJob {
		s := (float64(EndFlag) - float64(StartFlag)) / float64(UnixStep)
		return int(math.Ceil(s))
	}
	return 1
}

func nowStartEndTime() (string, string) {
	baseTime := time.Now().Unix()
	return strconv.FormatInt(baseTime, 10), strconv.FormatInt(baseTime+int64(UnixStep), 10)
}

func dayTime() string {
	start, _ := GetStartEndTime()
	s, _ := strconv.ParseInt(start, 10, 64)
	tm := time.Unix(s, 0).In(location)
	return tm.Format("20060102")
}
