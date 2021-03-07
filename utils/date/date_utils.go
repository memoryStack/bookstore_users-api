package date

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

/*
	on server save the time as UTC and on front-end covert that time acc. to the local time.
	this is the right way to handle the time in applications.
*/
