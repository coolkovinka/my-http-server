package handlers

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJoinPath(t *testing.T) {
	result, err := url.JoinPath("http://", "localhost:8080", "/hdfSDS")

	require.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/hdfSDS", result)
}
