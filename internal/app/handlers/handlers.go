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

func (h *Handlers) SetShortURL(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "wrong request type", http.StatusBadRequest)
		return
	}

	originalURL, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := h.storage.SetByOriginalURL(string(originalURL))

	response.WriteHeader(http.StatusCreated)
	_, err = response.Write([]byte(shortURL))
	if err != nil {
		http.Error(response, "failed writing response body", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) GetOriginalURL(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "wrong request type", http.StatusBadRequest)
		return
	}

	originalURL := h.storage.GetByURLPath(request.URL.Path)

	http.Redirect(response, request, originalURL, http.StatusTemporaryRedirect)

}
