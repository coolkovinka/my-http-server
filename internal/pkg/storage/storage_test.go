package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByShortURL(t *testing.T) {
	storage := NewStorage()

	tests := []struct {
		name        string
		path        string
		originalURL string
	}{
		{
			name:        "positive",
			path:        "/EwHXdJfB",
			originalURL: "https://practicum.yandex.ru/",
		},
		{
			name:        "negative",
			path:        "",
			originalURL: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			originalURL := storage.GetByURLPath(test.path)
			assert.Equal(t, test.originalURL, originalURL)
		})
	}
}

func TestSetByOriginalURL(t *testing.T) {
	storage := NewStorage()

	tests := []struct {
		name        string
		path        string
		originalURL string
	}{
		{
			name:        "positive",
			originalURL: "https://practicum.yandex.ru/",
			path:        "http://localhost:8080/EwHXdJfB",
		},
		{
			name:        "negative",
			originalURL: "",
			path:        "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := storage.SetByOriginalURL(test.originalURL)
			assert.Equal(t, test.path, path)
		})
	}
}
