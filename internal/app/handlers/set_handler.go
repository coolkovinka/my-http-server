package handlers

import (
	"io"
	"net/http"
)

func (h *Handlers) SetShortURL(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "method should be post", http.StatusBadRequest)
		return
	}

	originalURL, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "error reading request body", http.StatusInternalServerError)
		return
	}

	if string(originalURL) == "" {
		http.Error(response, "empty request body", http.StatusNotFound)
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
