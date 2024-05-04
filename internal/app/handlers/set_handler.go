package handlers

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func (h *Handlers) SetShortURL(response http.ResponseWriter, request *http.Request) {
	originalURL, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("error reading request body %v", err)
		http.Error(response, "something went wrong", http.StatusInternalServerError)
		return
	}

	if len(originalURL) == 0 {
		http.Error(response, "empty request body", http.StatusBadRequest)
		return
	}

	shortURL := h.storage.SetByOriginalURL(string(originalURL))

	scheme := "http://"
	baseURL, _ := url.JoinPath(scheme, h.cfg.ServerAddress, shortURL)

	response.WriteHeader(http.StatusCreated)

	_, err = response.Write([]byte(baseURL))
	if err != nil {
		log.Printf("error writing response body %v", err)
		http.Error(response, "something went wrong", http.StatusInternalServerError)
		return
	}
}
