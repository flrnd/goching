package goching

import (
	"fmt"
	"strconv"
	"strings"
)

type yarrow []string

type Hexagram struct {
	Number       int
	BinaryString string
}

type reading struct {
	Hexagram    Hexagram
	Lines       yarrow
	MovingLines []int
	RelatingHex Hexagram
}

var yarrows = yarrow{"OYin", "OYang", "OYang", "OYang", "Yang", "Yang", "Yang", "Yang", "Yang",
	"Yin", "Yin", "Yin", "Yin", "Yin", "Yin", "Yin",
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

func (y yarrow) shuffle() yarrow {
	size := len(y)
	dest := make(yarrow, size)
	perm := rng.Perm(size)
	for index := range y {
		dest[index] = y[perm[index]]
	}

	return dest
}

var NewYarrows = yarrows.shuffle()

func toBinary(hex []string) string {
	var sb strings.Builder

	for _, element := range hex {
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

func movingLines(hex []string) []int {
	var moving []int

	for i, element := range hex {
		switch element {
		case "OYin":
			moving = append(moving, i)
		case "OYang":
			moving = append(moving, i)
		}
	}
	return moving
}

func (y yarrow) newCast() []string {
	size := 6
	cast := make([]string, size)
	for index := range cast {
		position := rng.Int() % len(y)
		cast[index] = y[position]
	}
	return cast
}

func (y yarrow) CastReading() reading {
	cast := y.newCast()
	binaryString := toBinary(cast)

	hexNumber, err := binaryStringToHexagram(binaryString)

	if err != nil {
		panic(err)
	}

	hex := Hexagram{
		Number:       hexNumber,
		BinaryString: binaryString,
	}

	movingLines := movingLines(cast)
	var relatingHex Hexagram

	if len(movingLines) > 0 {
		relatingHex = hex.findRelatingHexagram(movingLines)
	}

	return reading{
		Hexagram:    hex,
		Lines:       cast,
		MovingLines: movingLines,
		RelatingHex: relatingHex,
	}
}
