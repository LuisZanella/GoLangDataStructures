package main

import (
	"net/http"
	"time"
)

func Racer(a string, b string) (winner string) {
	aTimeLapsed := measureResponseTime(a)
	bTimeLapsed := measureResponseTime(b)
	if aTimeLapsed < bTimeLapsed {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
