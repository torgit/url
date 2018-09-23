package main

import (
	"testing"
)

func TestHelper(t *testing.T) {
	t.Run("generateNewUrl: generate new url from original url", func(t *testing.T) {
		stubUrlStore := &stubUrlStore{}

		originalUrl := "taskworld.com/workspace/uploadfilename"
		expectedUrl := url{1, "taskworld.com/workspace/uploadfilename", "taskworld.com/MQ=="}
		generatedUrl, err := generateNewUrl(stubUrlStore, originalUrl)

		assertNoError(t, err)
		assertStringEquals(t, generatedUrl.originalUrl, expectedUrl.originalUrl)
		assertStringEquals(t, generatedUrl.shortUrl, expectedUrl.shortUrl)
	})

	t.Run("splitUrlDomain: test domain and path", func(t *testing.T) {
		url := "taskworld.com/MQ=="
		expectedDomain := "taskworld.com/"
		expectedPath := "MQ=="
		gotDomain, gotPath := splitUrlDomain(url)
		assertStringEquals(t, gotDomain, expectedDomain)
		assertStringEquals(t, gotPath, expectedPath)
	})

	t.Run("getIdFromPath: test path -> id conversion", func(t *testing.T) {
		path := "MQ=="
		expectedId := 1
		got, err := getIdFromPath(path)
		assertNoError(t, err)
		assertIntEquals(t, got, expectedId)
	})

	t.Run("getPathFromId: test id -> path conversion", func(t *testing.T) {
		id := 1
		expectedPath := "MQ=="
		got := getPathFromId(id)
		assertStringEquals(t, got, expectedPath)
	})
}
