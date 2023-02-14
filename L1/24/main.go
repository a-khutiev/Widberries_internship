// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func New(a int, b int) *Point {
	return &Point{x: a,
		y: b,
	}
}
func main() {
	A := New(1, 1)
	B := New(11, 1)
	fmt.Print(culc(A, B))
}
func culc(A *Point, B *Point) float64 {
	res := math.Pow((math.Pow(float64(B.x-A.x), 2) + math.Pow(float64(B.y-A.y), 2)), 0.5)
	return res
}
