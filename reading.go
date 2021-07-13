package goching

import (
	"fmt"
	"strconv"
	"strings"
)

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

func CastReading(c cast) Reading {
	binaryString := c.asBinarySeqString()

	hexNumber, err := binaryStringToHexagram(binaryString)

	if err != nil {
		panic(err)
	}

	hexagram := Hexagram{
		Number:       hexNumber,
		BinaryString: binaryString,
	}

	var relatingHex Hexagram

	if movingLines := c.getMovingLines(); len(movingLines) > 0 {
		relatingHex = hexagram.findRelatingHexagram(movingLines)
	}

	return Reading{
		Hexagram:    hexagram,
		Lines:       c,
		MovingLines: c.getMovingLines(),
		RelatingHex: relatingHex,
	}
}
