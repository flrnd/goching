package goching

import "github.com/flrnd/gorng"

type yarrows []string

var stalks = yarrows{"OYin", "OYang", "OYang", "OYang", "Yang", "Yang", "Yang", "Yang", "Yang", "Yin", "Yin", "Yin", "Yin", "Yin", "Yin", "Yin"}

func shuffle() yarrows {
	dest := make(yarrows, len(stalks))
	perm := gorng.Rng.Perm(len(stalks))

	for index := range stalks {
		dest[index] = stalks[perm[index]]
	}

	return dest
}

// New returns a new set of yarrow stalks
func New() readingCast {
	y := shuffle()
	size := 6
	cast := make([]string, size)

	for index := range cast {
		position := gorng.Rng.Int() % len(y)
		cast[index] = y[position]
	}
	return cast
}
