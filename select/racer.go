package racer

import (
	"fmt"
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

func RacerCh(urlA, urlB string) (urlWinner string, err error) {
	select {
	// the fastest received by the channel wins
	case <-pingUrl(urlA):
		return urlA, nil
	case <-pingUrl(urlB):
		return urlB, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}

func pingUrl(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
