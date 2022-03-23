package tcpclient

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestConn_RandomTokenUnique(t *testing.T) {
	src := rand.NewSource(time.Now().UnixNano())
	c := Conn{respRand: rand.New(src)}
	n := 10000
	tokens := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		tokens[c.randomToken()] = struct{}{}
	}
	assert.Len(t, tokens, n)
}
