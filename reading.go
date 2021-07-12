package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type yarrow []string

type hexagram struct {
	Number       int
	BinaryString string
}

type reading struct {
	Hexagram    hexagram
	Lines       yarrow
	MovingLines []int
	RelatingHex hexagram
}

var yarrows = yarrow{"OYin", "OYang", "OYang", "OYang", "Yang", "Yang", "Yang", "Yang", "Yang",
	"Yin", "Yin", "Yin", "Yin", "Yin", "Yin", "Yin",
}

func (hex hexagram) findRelatingHexagram(lines []int) hexagram {
	bs := strings.Split(hex.BinaryString, "")
	for _, line := range lines {
		num, _ := strconv.Atoi(bs[line])
		bs[line] = strconv.Itoa(num ^ 1)
	}
	relatingHex := hexagram{}
	relatingHex.BinaryString = strings.Join(bs, "")
	relatingHex.Number = HexBinaryStringToNumber(relatingHex.BinaryString)
	return relatingHex
}

func (y yarrow) shuffle() yarrow {
	size := len(y)
	dest := make(yarrow, size)
	perm := rand.Perm(size)
	for index := range y {
		dest[index] = y[perm[index]]
	}

	return dest
}

var newYarrows = yarrows.shuffle()

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
		position := rand.Int() % len(y)
		cast[index] = y[position]
	}
	return cast
}

func (y yarrow) castReading() reading {
	cast := y.newCast()
	binaryString := toBinary(cast)

	hex := hexagram{
		Number:       HexBinaryStringToNumber(binaryString),
		BinaryString: binaryString,
	}

	movingLines := movingLines(cast)
	var relatingHex hexagram

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

func (r reading) print() {
	fmt.Printf("Hexagram: %v\n", r.Hexagram.Number)
	if len(r.MovingLines) > 0 {
		fmt.Print("Lines: ")
		for _, line := range r.MovingLines {
			fmt.Printf("%v ", line+1)
		}
		fmt.Println()
		fmt.Printf("Relating: %v\n", r.RelatingHex.Number)
	}
}

// main function
func main() {
	rand.Seed(time.Now().UnixNano())
	newYarrows.castReading().print()
}
