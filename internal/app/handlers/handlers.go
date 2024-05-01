package handlers

import (
	"io"
	"net/http"
)

type Storage interface {
	GetByURLPath(string) string
	SetByOriginalURL(string) string
}

type Handlers struct {
	storage Storage
}

func NewHandlers(storage Storage) *Handlers {
	return &Handlers{
		storage: storage,
	}
}

func (s *Handlers) GetOriginalURL(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "wrong request type", http.StatusBadRequest)
		return
	}

	originalURL := s.storage.GetByURLPath(request.URL.Path)
	if originalURL != "" {
		http.Redirect(response, request, originalURL, http.StatusTemporaryRedirect)
	} else {
		http.Error(response, "invalid request data", http.StatusBadRequest)
		return
	}
}

func (s *Handlers) SetShortURL(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "wrong request type", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := s.storage.SetByOriginalURL(string(body))
	if shortURL == "" {
		http.Error(response, "such an URL does not exist", http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusCreated)

	_, err = response.Write([]byte(shortURL))
	if err != nil {
		http.Error(response, "failed writing response body", http.StatusInternalServerError)
		return
	}
}
