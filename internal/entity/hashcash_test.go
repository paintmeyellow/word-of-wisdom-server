package entity

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHashcash_Compute(t *testing.T) {
	hc, err := NewHashcash(nil)
	require.NoError(t, err)
	nonce, hash := hc.Compute()
	assert.Equal(t, hc.Nonce, 0)
	assert.True(t, nonce > 0)
	assert.NotEmpty(t, hash)
}

func TestHashcash_Validate(t *testing.T) {
	hc, err := NewHashcash(nil)
	require.NoError(t, err)
	nonce, _ := hc.Compute()

	t.Run("not_valid", func(t *testing.T) {
		assert.False(t, hc.Validate())
	})
	t.Run("valid", func(t *testing.T) {
		hc.Nonce = nonce
		assert.True(t, hc.Validate())
	})
}
