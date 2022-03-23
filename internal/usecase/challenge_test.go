package usecase

import (
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
	"word-of-wisdom/internal/entity"
)

func TestChallengeUseCase_Request(t *testing.T) {
	uc := ChallengeUseCase{}
	res, err := uc.Request(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestChallengeUseCase_Verify(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	t.Run("invalid_hash", func(t *testing.T) {
		//r := NewRepoMock(mc).ExistsMock.Return(true)
		uc := ChallengeUseCase{}
		hc, err := entity.NewHashcash(nil)
		require.NoError(t, err)
		_, err = uc.Verify(hc, nil)
		assert.ErrorIs(t, err, ErrInvalidHash)
	})

	t.Run("challenge_expired", func(t *testing.T) {
		uc := ChallengeUseCase{}
		hc := hashcash()
		hc.Nonce = 399555
		hc.Date = time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
		_, err := uc.Verify(hc, nil)
		assert.ErrorIs(t, err, ErrChallengeExpired)
	})
}

func hashcash() *entity.Hashcash {
	bits := big.NewInt(1)
	bits.Lsh(bits, uint(256-20))
	return &entity.Hashcash{
		Bits:     bits,
		Rand:     nil,
		Resource: nil,
		Date:     time.Now(),
	}
}
