package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   UrlService
}

func (mw loggingMiddleware) GetShortUrl(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetShort",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetShortUrl(s)
	return
}

func (mw loggingMiddleware) GetOriginalUrl(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetOriginal",
			"input", s,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetOriginalUrl(s)
	return
}
