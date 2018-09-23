package main

import (
	"testing"
)

type stubUrlStore struct {
	urls []url
}

func (s *stubUrlStore) getById(id int) (url, error) {
	for i := range s.urls {
		if s.urls[i].id == id {
			return s.urls[i], nil
		}
	}
	return url{}, errorFailToFind
}

func (s *stubUrlStore) getByOriginalUrl(originalUrl string) (url, error) {
	for i := range s.urls {
		if s.urls[i].originalUrl == originalUrl {
			return s.urls[i], nil
		}
	}
	return url{}, errorFailToFind
}

func (s *stubUrlStore) upsertUrl(u url) error {
	var appened []url
	existing, err := s.getByOriginalUrl(u.originalUrl)
	if err != nil {
		u.id = len(s.urls) + 1
		appened = append(s.urls, u)
	} else {
		existing.shortUrl = u.shortUrl
		for i := range s.urls {
			if s.urls[i].originalUrl == existing.originalUrl {
				appened = append(appened, existing)
			} else {
				appened = append(appened, s.urls[i])
			}
		}
	}
	s.urls = appened
	return nil
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("expect error but got none")
	}
}

func assertStringEquals(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertIntEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
