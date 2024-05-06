package storage

import (
	"math/rand"
	"sync"
)

type Storage struct {
	repo  map[string]string
	mutex sync.RWMutex
}

func NewStorage() *Storage {
	// key = short, value = original
	repo := make(map[string]string)

	return &Storage{
		repo: repo,
	}
}

func (s *Storage) GetByURLPath(shortURL string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.repo[shortURL]
}

func (s *Storage) SetByOriginalURL(originalURL string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	shortURL := generateRandomString(rand.Intn(randLen))
	s.repo[shortURL] = originalURL

	return shortURL
}

func generateRandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randLen = 20
)
