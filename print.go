package goching

import (
	"github.com/fatih/color"
)

func printHexagram(lines []string) {
	for i := len(lines); i >= 0; i-- {
		switch lines[i] {
		case "OYin":
			color.Red("- -")
		case "OYang":
			color.Red("---")
		case "Yin":
			color.White("- -")
		case "Yang":
			color.White("---")
		}
	}
}

func (r Reading) print() {
	printHexagram(r.Lines)
}
