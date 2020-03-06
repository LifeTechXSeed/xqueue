package util

import (
	"strconv"
	"time"
)

func SetInterval(someFunc func(), milliseconds int, async bool) chan bool {
	interval := time.Duration(milliseconds) * time.Millisecond

	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				if async {
					go someFunc()
				} else {
					someFunc()
				}
			case <-clear:
				ticker.Stop()
				return
			}
		}
	}()

	return clear
}

func ConvertTimeStampToTime(timestamp string) time.Time {
	if timestamp == "" || len(timestamp) <= 3 {
		return time.Time{}
	}

	goTimeStamp := timestamp[:len(timestamp)-3]
	timeToInt, err := strconv.ParseInt(goTimeStamp, 10, 64)
	if err != nil {
		return time.Time{}
	}

	result := time.Unix(timeToInt, 0)
	return result
}
