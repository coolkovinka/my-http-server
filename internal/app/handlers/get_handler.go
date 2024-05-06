package handlers

import (
	"net/http"
	"strings"
)

func (h *Handlers) GetOriginalURL(response http.ResponseWriter, request *http.Request) {
	originalURL := h.storage.GetByURLPath(strings.TrimPrefix(request.URL.Path, "/"))

	if originalURL == "" {
		http.Error(response, "URL does not exist", http.StatusNotFound)
		return
	}

	http.Redirect(response, request, originalURL, http.StatusTemporaryRedirect)
}
