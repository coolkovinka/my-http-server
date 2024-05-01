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

	http.Redirect(response, request, "https://practicum.yandex.ru/", http.StatusTemporaryRedirect)

}

func (s *Handlers) SetShortURL(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "wrong request type", http.StatusBadRequest)
		return
	}

	_, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "invalid request body", http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusCreated)

	_, err = response.Write([]byte("http://localhost:8080/EwHXdJfB"))
	if err != nil {
		http.Error(response, "failed writing response body", http.StatusInternalServerError)
		return
	}
}
