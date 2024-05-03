package handlers

import (
	"bytes"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"my-http-server/config"
	"my-http-server/internal/pkg/storage"
)

func TestHandlers_SetShortURL(t *testing.T) {
	type want struct {
		code    int
		target  string
		method  string
		reqBody string
	}

	repo := storage.NewStorage()
	handler := NewHandlers(&config.Config{}, repo)

	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test",
			want: want{
				code:    http.StatusCreated,
				target:  "/",
				method:  http.MethodPost,
				reqBody: "https://go.dev/",
			},
		},
		{
			name: "negative test. Invalid request URL",
			want: want{
				code:   http.StatusNotFound,
				target: "/",
				method: http.MethodPost,
			},
		},
		{
			name: "negative test. Wrong method type",
			want: want{
				code:   http.StatusBadRequest,
				target: "/",
				method: http.MethodGet,
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			reqBody := test.want.reqBody

			request := httptest.NewRequest(test.want.method, test.want.target, bytes.NewBuffer([]byte(reqBody)))
			w := httptest.NewRecorder()

			handler.SetShortURL(w, request)

			res := w.Result()

			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			err = res.Body.Close()
			require.NoError(t, err)

			assert.Equal(t, test.want.code, res.StatusCode)
			assert.NotEmpty(t, string(resBody))
		})
	}
}
