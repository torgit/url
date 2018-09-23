package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           UrlService
}

func (mw instrumentingMiddleware) GetShortUrl(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetShortUrl", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetShortUrl(s)
	return
}

func (mw instrumentingMiddleware) GetOriginalUrl(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetOriginalUrl", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetOriginalUrl(s)
	return
}
