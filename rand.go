/*
 * Check: https://blog.gopheracademy.com/advent-2017/a-tale-of-two-rands/
 */

package goching

import (
	crand "crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

var rng = randInit()

type cryptoSource struct{}

func (s *cryptoSource) Seed(seed int64) {}

func (s *cryptoSource) Uint64() (value uint64) {
	binary.Read(crand.Reader, binary.BigEndian, &value)
	return value
}

func (s *cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func randInit() (myRand *mrand.Rand) {
	var src mrand.Source64 = &cryptoSource{}
	myRand = mrand.New(src)

	return myRand
}
