package main

import (
	"strconv"
	"strings"
)

func dateToHour(date string) string {
	dateSplit := strings.Split(date, " ")
	time := dateSplit[3]
	return time
}
func getSecFromDate(date string) string {
	//Fri Apr 28 17:25:22 +0000 2017
	dateSplit := strings.Split(date, " ")
	time := dateSplit[3]
	sec := strings.Split(time, ":")[2]
	return sec
}
func anonymousName(originalName string, date string) string {
	s := strings.Split(originalName, "")
	//first letter and the last-1 (prelast / penúltima) letter + seconds
	generatedName := s[0] + s[len(s)-2] + strconv.Itoa(len(s))
	return generatedName
}
