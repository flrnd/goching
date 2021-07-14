package goching

type yarrows []string

type cast []string

// Hexagram number and a binary sequence string
type Hexagram struct {
	Number       int
	BinaryString string
}

// Reading is an I Ching reading cast
type Reading struct {
	Hexagram    Hexagram
	Lines       cast
	MovingLines []int
	RelatingHex Hexagram
}
