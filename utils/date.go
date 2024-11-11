package utils

import "time"

func GetNowDate()time.Time{
	t := GetNowTime()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return newTime
}