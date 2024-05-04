package handlers

import "net/http"

func (h *Handlers) GetOriginalURL(response http.ResponseWriter, request *http.Request) {
	originalURL := h.storage.GetByURLPath(request.URL.Path)
	if originalURL == "" {
		http.Error(response, "URL does not exist", http.StatusNotFound)
		return
	}

	http.Redirect(response, request, originalURL, http.StatusTemporaryRedirect)
}
