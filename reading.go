package goching

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func isLine(line string, pattern string) bool {
	var validLine = regexp.MustCompile(fmt.Sprintf("(?i)%s$", pattern))
	return validLine.MatchString(line)
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

func (c readingCast) asBinarySeqString() string {
	var sb strings.Builder

	for _, line := range c {
		if isLine(line, "yin") {
			fmt.Fprintf(&sb, "0")
		}
		if isLine(line, "yang") {
			fmt.Fprintf(&sb, "1")
		}
	}

	return sb.String()
}

func (c readingCast) getMovingLines() []int {
	var lines []int

	for i, line := range c {
		if isLine(line, "oyin") || isLine(line, "oyang") {
			lines = append(lines, i)
		}
	}

	return lines
}

// CastReading returns a full formed Reading struct
func CastReading(c readingCast) Reading {
	binaryString := c.asBinarySeqString()

	hexNumber, err := binaryStringToHexagram(binaryString)

	if err != nil {
		panic(err)
	}

	hexagram := Hexagram{
		Number:       hexNumber,
		BinaryString: binaryString,
	}

	var relating Hexagram

	if lines := c.getMovingLines(); len(lines) > 0 {
		relating = hexagram.findRelatingHexagram(lines)
	}

	return Reading{
		Hexagram: hexagram,
		Lines:    c.getMovingLines(),
		Relating: relating,
	}
}
