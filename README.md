# I Ching go

I Ching library written in Go

## Example

```Go
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
  // NewCast will return an slice with 6 lines
  newCast := goching.NewCast

  // CastReading will return the Reading struct
  // from your 6 line slice
  myReading := goching.CastReading(NewCast)
  print(myReading)
}
```

[![Go Report Card](https://goreportcard.com/badge/github.com/flrnd/goching)](https://goreportcard.com/report/github.com/flrnd/goching)
