package goching

type yarrows []string

type readingCast []string

// Hexagram number and a binary sequence string
type Hexagram struct {
	Number       int
	BinaryString string
}

// Reading is an I Ching reading cast
type Reading struct {
	Hexagram Hexagram
	Relating Hexagram
	Lines    []int
}
