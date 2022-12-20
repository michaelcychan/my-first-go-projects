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

var timeOutTenSecond = 10 * time.Second

func RacerCh(urlA, urlB string) (urlWinner string, err error) {
	// refactor the timeout configurable out
	return RacerConfigurable(urlA, urlB, timeOutTenSecond)
}

func RacerConfigurable(urlA, urlB string, timeout time.Duration) (urlWinner string, err error) {
	select {
	// the fastest received by the channel wins
	case <-pingUrl(urlA):
		return urlA, nil
	case <-pingUrl(urlB):
		return urlB, nil
	case <-time.After(timeout):
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
