package goching

var stalks = yarrows{"OYin", "OYang", "OYang", "OYang", "Yang", "Yang", "Yang", "Yang", "Yang",
	"Yin", "Yin", "Yin", "Yin", "Yin", "Yin", "Yin",
}

func (y yarrows) shuffle() yarrows {
	size := len(y)
	dest := make(yarrows, size)
	perm := rng.Perm(size)
	for index := range y {
		dest[index] = y[perm[index]]
	}

	return dest
}

func (y yarrows) getLines() cast {
	size := 6
	cast := make([]string, size)
	for index := range cast {
		position := rng.Int() % len(y)
		cast[index] = y[position]
	}
	return cast
}

var NewCast = stalks.shuffle().getLines()
