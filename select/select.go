package _select

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
	//aDuration := measureResponsetime(a)
	//bDuration := measureResponsetime(b)
	//
	//if aDuration < bDuration {
	//	return a
	//}
	//
	//return b
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

//
//func measureResponsetime(url string) time.Duration {
//	start := time.Now()
//	http.Get(url)
//	return time.Since(start)
//}
