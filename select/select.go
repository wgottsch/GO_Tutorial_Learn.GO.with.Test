package _select

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponsetime(a)
	bDuration := measureResponsetime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponsetime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)

}
