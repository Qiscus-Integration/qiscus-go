package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	qiscusAppID     = "test-qiscus-app-id"
	qiscusSecretKey = "test-qiscus-secret-key"
)

func TestNewSDK(t *testing.T) {
	c := NewSDK(qiscusAppID, qiscusSecretKey)
	assert.Equal(t, c.QiscusAppID(), qiscusAppID)
	assert.Equal(t, c.QiscusSecretKey(), qiscusSecretKey)
}
