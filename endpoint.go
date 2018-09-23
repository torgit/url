package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetShortUrlEndpoint(svc UrlService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getUrlRequest)
		v, err := svc.GetShortUrl(req.Url)
		if err != nil {
			return getUrlResponse{v, err.Error()}, nil
		}
		return getUrlResponse{v, ""}, nil
	}
}

func makeGetOriginalUrlEndpoint(svc UrlService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getUrlRequest)
		v, err := svc.GetOriginalUrl(req.Url)
		if err != nil {
			return getUrlResponse{v, err.Error()}, nil
		}
		return getUrlResponse{v, ""}, nil
	}
}
