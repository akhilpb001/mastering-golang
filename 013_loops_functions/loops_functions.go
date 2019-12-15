package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	fmt.Println(sqrt(2))
}
