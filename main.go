package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "urlshortener",
		Subsystem: "url_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "urlshortener",
		Subsystem: "url_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	urlStore, err := newDbStore()
	if err != nil {
		panic(err)
	}

	var svc UrlService
	svc = newUrlService(urlStore)
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, svc}

	getShortHandler := httptransport.NewServer(
		makeGetShortUrlEndpoint(svc),
		decodeGetUrlRequest,
		encodeGetUrlResponse,
	)

	getOriginalHandler := httptransport.NewServer(
		makeGetOriginalUrlEndpoint(svc),
		decodeGetUrlRequest,
		encodeGetUrlResponse,
	)

	http.Handle("/getShortUrl", getShortHandler)
	http.Handle("/getOriginalUrl", getOriginalHandler)

	http.Handle("/metrics", promhttp.Handler())

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
