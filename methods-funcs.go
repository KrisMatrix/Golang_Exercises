package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(b Vertex) float64 {
	return math.Sqrt(b.X*b.X + b.Y*b.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
