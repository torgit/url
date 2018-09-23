package main

import (
	"testing"
)

func TestService(t *testing.T) {
	t.Run("GetShortUrl: create new record from original url", func(t *testing.T) {
		stubUrlStore := &stubUrlStore{}
		service := urlService{stubUrlStore}

		originalUrl := "taskworld.com/workspace/uploadfilename"
		expectedShortUrl := "taskworld.com/MQ=="

		gotShortUrl, err := service.GetShortUrl(originalUrl)

		assertNoError(t, err)
		assertStringEquals(t, gotShortUrl, expectedShortUrl)
		assertStringEquals(t, gotShortUrl, expectedShortUrl)
	})

	t.Run("GetOriginalUrl: get existing original url from short url", func(t *testing.T) {
		stubUrlStore := &stubUrlStore{}
		service := urlService{stubUrlStore}

		shortUrl := "taskworld.com/MQ=="
		expectedOriginalUrl := "taskworld.com/workspace/uploadfilename"

		//Insert record before searching
		_, _ = service.GetShortUrl(expectedOriginalUrl)

		gotOriginalUrl, err := service.GetOriginalUrl(shortUrl)

		assertNoError(t, err)
		assertStringEquals(t, gotOriginalUrl, expectedOriginalUrl)
		assertStringEquals(t, gotOriginalUrl, expectedOriginalUrl)
	})

	t.Run("GetOriginalUrl: cannot get original url from non existing record", func(t *testing.T) {
		stubUrlStore := &stubUrlStore{}
		service := urlService{stubUrlStore}

		nonExistingShortUrl := "taskworld.com/Mg=="
		originalUrl := "taskworld.com/workspace/uploadfilename"

		//Insert record before searching
		_, _ = service.GetShortUrl(originalUrl)

		_, err := service.GetOriginalUrl(nonExistingShortUrl)

		assertError(t, err)
	})
}
