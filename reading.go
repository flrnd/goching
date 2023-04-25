package goching

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/flrnd/goching/dictionary"
)

type readingCast []string

// Hexagram number and a binary sequence string
type Hexagram struct {
	Number       *int
	BinaryString string
}

// Reading is an I Ching reading cast
type Reading struct {
	Hexagram *Hexagram
	Relating *Hexagram
	Lines    []int
}

func isLine(line string, pattern string) bool {
	var validLine = regexp.MustCompile(fmt.Sprintf("(?i)%s$", pattern))
	return validLine.MatchString(line)
}

func (hex Hexagram) findRelatingHexagram(lines []int) *Hexagram {
	bs := strings.Split(hex.BinaryString, "")
	for _, line := range lines {
		num, _ := strconv.Atoi(bs[line])
		bs[line] = strconv.Itoa(num ^ 1)
	}
	relatingHex := Hexagram{}
	relatingHex.BinaryString = strings.Join(bs, "")
	relating, err := dictionary.GetHexagram(relatingHex.BinaryString)

	if err != nil && errors.Is(err, dictionary.ErrInvalidBinaryString) {
		return nil
	}

	relatingHex.Number = relating
	return &relatingHex
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

// CastReading returns a Reading
func CastReading(c readingCast) *Reading {
	binaryString := c.asBinarySeqString()
	lines := c.getMovingLines()

	hexNumber, err := dictionary.GetHexagram(binaryString)
	if err != nil && errors.Is(err, dictionary.ErrInvalidBinaryString) {
		log.Printf("CastReading error: %v\n", err)
		return nil
	}

	hexagram := &Hexagram{
		Number:       hexNumber,
		BinaryString: binaryString,
	}

	var relating *Hexagram

	if len(lines) > 0 {
		relating = hexagram.findRelatingHexagram(lines)
	}

	return &Reading{
		Hexagram: hexagram,
		Lines:    lines,
		Relating: relating,
	}
}
