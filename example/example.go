package main

import (
	"fmt"

	"github.com/flrnd/goching"
)

func print(reading goching.Reading) {
	fmt.Printf("Hexagram: %v\n", reading.Hexagram.Number)
	if len(reading.MovingLines) > 0 {
		fmt.Print("Lines: ")
		for _, line := range reading.MovingLines {
			fmt.Printf("%v ", line+1)
		}
		fmt.Println()
		fmt.Printf("Relating: %v\n", reading.RelatingHex.Number)
	}
}

func main() {
	myReading := goching.CastReading(goching.NewCast)

	print(myReading)
}
