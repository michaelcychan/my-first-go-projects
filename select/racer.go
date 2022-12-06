package racer

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (urlWinner string) {
	aDuration := measureResponseTime(urlA)
	bDuration := measureResponseTime(urlB)

	if aDuration < bDuration {
		return urlA
	} else {
		return urlB
	}
}

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}
