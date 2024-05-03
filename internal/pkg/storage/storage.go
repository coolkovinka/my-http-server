package storage

import (
	"math/rand"
	"sync"
	"time"
)

type Storage struct {
	repo  map[string]string
	mutex sync.RWMutex
}

func NewStorage() *Storage {
	// key = original, value = short
	repo := make(map[string]string)

	return &Storage{
		repo: repo,
	}
}

func (s *Storage) GetByURLPath(shortURL string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for k, v := range s.repo {
		if v == shortURL {
			return k
		}
	}

	return ""
}

func (s *Storage) SetByOriginalURL(originalURL string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	shortURL := `/` + generateRandomString()
	s.repo[originalURL] = shortURL

	return shortURL
}

func generateRandomString() string {
	length := 8
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}
