package storage

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomString(t *testing.T) {
	randStrLen := rand.Intn(20)

	assert.Len(t, generateRandomString(randStrLen), randStrLen)
}
