package goching

type yarrows []string

type cast []string

type Hexagram struct {
	Number       int
	BinaryString string
}

type Reading struct {
	Hexagram    Hexagram
	Lines       cast
	MovingLines []int
	RelatingHex Hexagram
}
