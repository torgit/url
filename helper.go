package main

import (
	"encoding/base64"
	"regexp"
	"strconv"
)

func generateNewUrl(urlStore urlStore, originalUrl string) (url, error) {
	newUrl := url{originalUrl: originalUrl}
	err := urlStore.upsertUrl(newUrl)
	if err != nil {
		return url{}, nil
	}
	newUrl, err = urlStore.getByOriginalUrl(originalUrl)
	shortUrlPath := getPathFromId(newUrl.id)
	domain, _ := splitUrlDomain(originalUrl)
	newUrl.shortUrl = domain + shortUrlPath
	return newUrl, nil
}

func splitUrlDomain(s string) (string, string) {
	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)/?`)
	domain := re.FindString(s)
	path := re.ReplaceAllString(s, "")
	return domain, path
}

func getIdFromPath(path string) (int, error) {
	decoded, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		return 0, errorInvalidRequestingUrl
	}
	id, err := strconv.Atoi(string(decoded))
	if err != nil {
		return 0, errorInvalidRequestingUrl
	}
	return id, nil
}

func getPathFromId(id int) string {
	idStr := strconv.Itoa(id)
	encoded := base64.StdEncoding.EncodeToString([]byte(idStr))
	return encoded
}
