package storage

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) GetByURLPath(path string) string {
	for k, v := range data {
		if v == path {
			return k
		}
	}

	return ""
}

func (s *Storage) SetByOriginalURL(originalURL string) string {
	if val, ok := data[originalURL]; ok {
		return "http://localhost:8080" + val
	}

	return ""
}

// key is originalURL, value is shortURL
var data = map[string]string{
	"https://practicum.yandex.ru/":         "/EwHXdJfB",
	"https://practicum.yandex.ru/catalog/": "/IKsdfRWf",
	"https://practicum.yandex.ru/profile/": "/jFGkFRTs",
}
