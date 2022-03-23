package entity

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math"
	"math/big"
	"strconv"
	"time"
)

const (
	targetBits           = 20
	bytesToRead   int    = 8
	maxIterations int    = math.MaxInt64
	timeFormat    string = "060102150405" // YYMMDDhhmmss
)

type Hashcash struct {
	Bits     *big.Int
	Date     time.Time
	Resource []byte
	Rand     []byte
	Nonce    int
}

func NewHashcash(resource []byte) (*Hashcash, error) {
	bits := big.NewInt(1)
	bits.Lsh(bits, uint(256-targetBits))
	b, err := randomBytes(bytesToRead)
	if err != nil {
		return nil, err
	}
	pow := Hashcash{
		Bits:     bits,
		Rand:     b,
		Resource: resource,
		Date:     time.Now(),
	}
	return &pow, nil
}

func (hc *Hashcash) Compute() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	for nonce < maxIterations {
		data := hc.prepareData(nonce)
		hash = sha256.Sum256(data)
		//fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(hc.Bits) == -1 {
			break
		} else {
			nonce++
		}
	}
	//fmt.Println()
	return nonce, hash[:]
}

func (hc *Hashcash) Validate() bool {
	var hashInt big.Int
	hash := hc.Hash()
	hashInt.SetBytes(hash)
	return hashInt.Cmp(hc.Bits) == -1
}

func (hc *Hashcash) Hash() []byte {
	data := hc.prepareData(hc.Nonce)
	hash := sha256.Sum256(data)
	return hash[:]
}

func (hc *Hashcash) HashToString() string {
	return hex.EncodeToString(hc.Hash())
}

func (hc *Hashcash) prepareData(nonce int) []byte {
	return bytes.Join([][]byte{
		[]byte(strconv.Itoa(targetBits)),
		[]byte(hc.Date.Format(timeFormat)),
		hc.Resource,
		hc.Rand,
		[]byte(strconv.Itoa(nonce)),
	}, []byte{})
}

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}
