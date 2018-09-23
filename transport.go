package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type getUrlRequest struct {
	Url string `json:"url"`
}

type getUrlResponse struct {
	Url string `json:"url"`
	Err string `json:"err,omitempty`
}

func decodeGetUrlRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeGetUrlResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
