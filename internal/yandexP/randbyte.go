package yandexp

import (
	"encoding/binary"
	"io"
	"math/rand"
	"strings"
)

type generator struct {
	rnd rand.Source
}

// Read implements io.Reader.
func (g *generator) Read(bytes []byte) (n int, err error) {
	// old desision
	// for i := range bytes {
	// 	randInt := g.rnd.Int63()
	// 	randByte := byte(randInt)
	// 	bytes[i] = randByte
	// }

	for i := 0; i+8 < len(bytes); i += 8 {
		j := i + 8
		binary.LittleEndian.PutUint64(bytes[i:j], uint64(g.rnd.Int63()))
	}
	return len(bytes), nil
}

func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

type Hasher interface {
	io.Writer
	Hash() byte
}

type hash struct {
	result byte
}

func NewHasher(_init byte) Hasher {
	return &hash{
		result: _init,
	}
}

func (h *hash) Write(bytes []byte) (n int, err error) {
	for _, b := range bytes {
		h.result = (h.result^b)<<1 + b%2
	}
	return len(bytes), nil
}

func (h hash) Hash() byte {
	return h.result
}

func Mul(a any, b int) (d any) {
	switch s := a.(type) {
	case string:
		d = strings.Repeat(s, b)
	case int:
		d = s * b
	default:
		d = "Неизвестный тип"
	}
	return d
}
