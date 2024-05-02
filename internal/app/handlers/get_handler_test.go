package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"my-http-server/config"
	"my-http-server/internal/pkg/storage"
)

func TestHandlers_GetOriginalURL(t *testing.T) {
	type want struct {
		code   int
		target string
		method string
	}

	cfg := config.Config{
		ServerAddress: "localhost:8080",
		ServerBaseURL: "http://localhost:8080",
	}

	originalURL := "https://go.dev/"

	repo := storage.NewStorage()
	shortURL := repo.SetByOriginalURL(originalURL)
	handler := NewHandlers(&cfg, repo)

	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test. Valid request short URL",
			want: want{
				code:   http.StatusTemporaryRedirect,
				target: shortURL,
				method: http.MethodGet,
			},
		},
		{
			name: "negative test. Invalid request short URL",
			want: want{
				code:   http.StatusNotFound,
				target: "/blabla",
				method: http.MethodGet,
			},
		},
		{
			name: "negative test. Wrong method type",
			want: want{
				code:   http.StatusBadRequest,
				target: "/",
				method: http.MethodPost,
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			request := httptest.NewRequest(test.want.method, test.want.target, nil)
			w := httptest.NewRecorder()
			handler.GetOriginalURL(w, request)

			res := w.Result()
			err := res.Body.Close()
			require.NoError(t, err)

			assert.Equal(t, test.want.code, res.StatusCode)
		})
	}
}
