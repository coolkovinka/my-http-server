package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomString(t *testing.T) {
	assert.Len(t, generateRandomString(), 8)
}
