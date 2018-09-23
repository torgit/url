package main

import (
	"fmt"
)

type UrlService interface {
	GetShortUrl(string) (string, error)
	GetOriginalUrl(string) (string, error)
}

type urlService struct {
	urlStore urlStore
}

func newUrlService(store urlStore) *urlService {
	return &urlService{store}
}

var errorInvalidRequestingUrl = fmt.Errorf("invalid requesting url")

func (service *urlService) GetShortUrl(originalUrl string) (string, error) {
	url, err := service.urlStore.getByOriginalUrl(originalUrl)
	if err == nil {
		return url.shortUrl, nil
	}
	newUrl, err := generateNewUrl(service.urlStore, originalUrl)
	if err != nil {
		return "", err
	}
	err = service.urlStore.upsertUrl(newUrl)
	if err != nil {
		return "", err
	}
	return newUrl.shortUrl, nil
}

func (service *urlService) GetOriginalUrl(shortUrl string) (string, error) {
	_, path := splitUrlDomain(shortUrl)
	id, err := getIdFromPath(path)
	if err != nil {
		return "", err
	}
	u, err := service.urlStore.getById(id)
	if err != nil {
		return "", err
	}
	return u.originalUrl, nil
}
