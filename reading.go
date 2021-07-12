package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type hexagram struct {
	Number       int
	BinaryString string
}

type reading struct {
	Hexagram    hexagram
	Lines       []string
	MovingLines []int
	RelatingHex hexagram
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

func shuffle(src []string) []string {
	size := len(src)
	dest := make([]string, size)
	perm := rand.Perm(size)
	for index := range src {
		dest[index] = src[perm[index]]
	}

	return dest
}

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

func generateCast(yarrow []string) []string {
	size := 6
	cast := make([]string, size)
	for index := range cast {
		position := rand.Int() % len(yarrow)
		cast[index] = yarrow[position]
	}
	return cast
}

func generateReading(yarrow []string) reading {
	cast := generateCast(yarrow)
	hex := hexagram{}
	relatingHex := hexagram{}
	reading := reading{
		Hexagram:    hexagram{},
		Lines:       []string{},
		MovingLines: []int{},
		RelatingHex: hexagram{},
	}

	hex.BinaryString = toBinary(cast)
	hex.Number = HexBinaryStringToNumber(hex.BinaryString)

	movingLines := movingLines(cast)

	reading.Hexagram = hex
	reading.Lines = cast
	reading.MovingLines = movingLines
	if len(movingLines) > 0 {
		relatingHex = reading.Hexagram.findRelatingHexagram(movingLines)
	}
	reading.RelatingHex = relatingHex
	return reading
}

// main function
func main() {
	rand.Seed(time.Now().UnixNano())

	yarrow := []string{"OYin", "OYang", "OYang", "OYang", "Yang", "Yang", "Yang", "Yang", "Yang",
		"Yin", "Yin", "Yin", "Yin", "Yin", "Yin", "Yin",
	}

	shuffled := shuffle(yarrow)

	newReading := generateReading(shuffled)

	println(newReading.Hexagram.Number)
}
