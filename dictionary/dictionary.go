package dictionary

import (
	"errors"
	"regexp"
)

// dictNumberToString map with hex number and its binary string
var dictNumberToString = map[int]string{
	1:  "111111",
	2:  "000000",
	3:  "100010",
	4:  "010001",
	5:  "111010",
	6:  "010111",
	7:  "010000",
	8:  "000010",
	9:  "111011",
	10: "110111",
	11: "111000",
	12: "000111",
	13: "101111",
	14: "111101",
	15: "001000",
	16: "000100",
	17: "100110",
	18: "011001",
	19: "110000",
	20: "000011",
	21: "100101",
	22: "101001",
	23: "000001",
	24: "100000",
	25: "100111",
	26: "111001",
	27: "100001",
	28: "011110",
	29: "010010",
	30: "101101",
	31: "001110",
	32: "011100",
	33: "001111",
	34: "111100",
	35: "000101",
	36: "101000",
	37: "101011",
	38: "110101",
	39: "001010",
	40: "010100",
	41: "110001",
	42: "100011",
	43: "111110",
	44: "011111",
	45: "000110",
	46: "011000",
	47: "010110",
	48: "011010",
	49: "101110",
	50: "011101",
	51: "100100",
	52: "001001",
	53: "001011",
	54: "110100",
	55: "101100",
	56: "001101",
	57: "011011",
	58: "110110",
	59: "010011",
	60: "110010",
	61: "110011",
	62: "001100",
	63: "101010",
	64: "010101",
}

// dictStringToNumber is a map with binary strings and hex number
var dictStringToNumber = map[string]int{
	"111111": 1,
	"000000": 2,
	"100010": 3,
	"010001": 4,
	"111010": 5,
	"010111": 6,
	"010000": 7,
	"000010": 8,
	"111011": 9,
	"110111": 10,
	"111000": 11,
	"000111": 12,
	"101111": 13,
	"111101": 14,
	"001000": 15,
	"000100": 16,
	"100110": 17,
	"011001": 18,
	"110000": 19,
	"000011": 20,
	"100101": 21,
	"101001": 22,
	"000001": 23,
	"100000": 24,
	"100111": 25,
	"111001": 26,
	"100001": 27,
	"011110": 28,
	"010010": 29,
	"101101": 30,
	"001110": 31,
	"011100": 32,
	"001111": 33,
	"111100": 34,
	"000101": 35,
	"101000": 36,
	"101011": 37,
	"110101": 38,
	"001010": 39,
	"010100": 40,
	"110001": 41,
	"100011": 42,
	"111110": 43,
	"011111": 44,
	"000110": 45,
	"011000": 46,
	"010110": 47,
	"011010": 48,
	"101110": 49,
	"011101": 50,
	"100100": 51,
	"001001": 52,
	"001011": 53,
	"110100": 54,
	"101100": 55,
	"001101": 56,
	"011011": 57,
	"110110": 58,
	"010011": 59,
	"110010": 60,
	"110011": 61,
	"001100": 62,
	"101010": 63,
	"010101": 64,
}

// ErrInvalidHexagramNumber is returned when request an invalid hexagram
var ErrInvalidHexagramNumber = errors.New("Invalid hexagram number")

// ErrInvalidBinaryString is returnes with invalid binary string
var ErrInvalidBinaryString = errors.New("Invalid binary string")

// GetBinaryString returns a binary string from an hexagram
func GetBinaryString(hexagram int) (b *string, e error) {
	if !isValidHexagram(hexagram) {
		return nil, ErrInvalidHexagramNumber
	}

	bString := dictNumberToString[hexagram]
	return &bString, nil
}

// GetHexagram returns the hexagram from a binary string
func GetHexagram(binaryString string) (h *int, e error) {
	if !isValidBinaryString(binaryString) {
		return nil, ErrInvalidBinaryString
	}

	hex := dictStringToNumber[binaryString]
	return &hex, nil
}

func isValidHexagram(h int) bool {
	return h > 0 && h < 65
}

func isValidBinaryString(s string) bool {
	match, _ := regexp.MatchString(`[01]{6}`, s)
	return match
}
