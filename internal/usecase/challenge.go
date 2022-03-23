package usecase

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
	"word-of-wisdom/internal/data"
	"word-of-wisdom/internal/entity"
)

var (
	ErrInvalidResource  = errors.New("err invalid resource")
	ErrInvalidHash      = errors.New("err invalid hash")
	ErrChallengeExpired = errors.New("err challenge expired")
	ErrCreateHashcash   = errors.New("err create hashcash")
)

type Repo interface {
	Store(hash string)
	Exists(hash string) bool
}

type ChallengeUseCase struct {
	Repo Repo
}

func (uc *ChallengeUseCase) Request(resource []byte) ([]byte, error) {
	hc, err := entity.NewHashcash(resource)
	if err != nil {
		return nil, ErrCreateHashcash
	}
	return json.Marshal(hc)
}

func (uc *ChallengeUseCase) Verify(hc *entity.Hashcash, resource []byte) (string, error) {
	if !hc.Validate() {
		return "", ErrInvalidHash
	}
	if time.Now().Sub(hc.Date) > time.Minute {
		return "", ErrChallengeExpired
	}
	if !bytes.Equal(hc.Resource, resource) {
		return "", ErrInvalidResource
	}
	hash := hc.HashToString()
	if uc.Repo.Exists(hash) {
		return "", ErrInvalidHash
	}
	uc.Repo.Store(hash)
	return data.RandomQuote(), nil
}
