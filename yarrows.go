package goching

import "github.com/flrnd/gorng"

var stalks = yarrows{"OYin", "OYang", "OYang", "OYang", "Yang", "Yang", "Yang", "Yang", "Yang",
	"Yin", "Yin", "Yin", "Yin", "Yin", "Yin", "Yin",
}

func (y yarrows) shuffle() yarrows {
	size := len(y)
	dest := make(yarrows, size)
	perm := gorng.Rng.Perm(size)
	for index := range y {
		dest[index] = y[perm[index]]
	}

	return dest
}

func (y yarrows) getLines() cast {
	size := 6
	cast := make([]string, size)
	for index := range cast {
		position := gorng.Rng.Int() % len(y)
		cast[index] = y[position]
	}
	return cast
}

// NewCast returns a slice with 6 string cast lines
var NewCast = stalks.shuffle().getLines()
