package main

import (
	"fmt"

	"github.com/flrnd/goching"
)

func print(reading goching.Reading) {
	fmt.Printf("Hexagram: %v\n", reading.Hexagram.Number)
	if len(reading.Lines) > 0 {
		fmt.Print("Lines: ")
		for _, line := range reading.Lines {
			fmt.Printf("%v ", line+1)
		}
		fmt.Println()
		fmt.Printf("Relating: %v\n", reading.Relating.Number)
	}
}

func main() {
	myReading := goching.CastReading(goching.NewStalks())

	print(*myReading)
}
