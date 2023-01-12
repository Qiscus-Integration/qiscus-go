package multichannel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	qiscusAppID     = "test-qiscus-app-id"
	qiscusSecretKey = "test-qiscus-secret-key"
)

func TestNewMultichannel(t *testing.T) {
	c := NewMultichannel(qiscusAppID, qiscusSecretKey)

	assert.Equal(t, c.QiscusAppID(), qiscusAppID)
	assert.Equal(t, c.QiscusSecretKey(), qiscusSecretKey)
}
