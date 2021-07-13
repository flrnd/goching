package goching

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (hex Hexagram) findRelatingHexagram(lines []int) Hexagram {
	bs := strings.Split(hex.BinaryString, "")
	for _, line := range lines {
		num, _ := strconv.Atoi(bs[line])
		bs[line] = strconv.Itoa(num ^ 1)
	}
	relatingHex := Hexagram{}
	relatingHex.BinaryString = strings.Join(bs, "")
	relating, _ := binaryStringToHexagram(relatingHex.BinaryString)
	relatingHex.Number = relating
	return relatingHex
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

func (c cast) asBinarySeqString() string {
	var sb strings.Builder

	for _, element := range c {
		switch element {
		case "OYin":
			fmt.Fprintf(&sb, "0")
		case "OYang":
			fmt.Fprintf(&sb, "1")
		case "Yin":
			fmt.Fprintf(&sb, "0")
		case "Yang":
			fmt.Fprintf(&sb, "1")
		}
	}

	return sb.String()
}

func (c cast) getMovingLines() []int {
	var lines []int

	for i, line := range c {
		switch line {
		case "OYin":
			lines = append(lines, i)
		case "OYang":
			lines = append(lines, i)
		}
	}
	return lines
}

func (y yarrows) newCast() cast {
	size := 6
	cast := make([]string, size)
	for index := range cast {
		position := rng.Int() % len(y)
		cast[index] = y[position]
	}
	return cast
}

func (y yarrows) CastReading() Reading {
	cast := y.newCast()
	binaryString := cast.asBinarySeqString()

	hexNumber, err := binaryStringToHexagram(binaryString)

	if err != nil {
		panic(err)
	}

	hex := Hexagram{
		Number:       hexNumber,
		BinaryString: binaryString,
	}

	var relatingHex Hexagram

	if movingLines := cast.getMovingLines(); len(movingLines) > 0 {
		relatingHex = hex.findRelatingHexagram(movingLines)
	}

	return Reading{
		Hexagram:    hex,
		Lines:       cast,
		MovingLines: cast.getMovingLines(),
		RelatingHex: relatingHex,
	}
}
