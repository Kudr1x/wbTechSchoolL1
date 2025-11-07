package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.x, p.y)
}

func main() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	fmt.Printf("Точка 1: %s\n", point1)
	fmt.Printf("Точка 2: %s\n\n", point2)

	dist1 := Distance(point1, point2)
	fmt.Printf("Расстояние: %.4f\n", dist1)

	p1 := NewPoint(0, 0)
	p2 := NewPoint(5, 0)
	fmt.Printf("Расстояние между %s и %s: %.2f\n", p1, p2, Distance(p1, p2))

	p3 := NewPoint(0, 0)
	p4 := NewPoint(0, 3)
	fmt.Printf("Расстояние между %s и %s: %.2f\n", p3, p4, Distance(p3, p4))

	p5 := NewPoint(0, 0)
	p6 := NewPoint(1, 1)
	fmt.Printf("Расстояние между %s и %s: %.4f\n", p5, p6, Distance(p5, p6))

	p7 := NewPoint(-3, -4)
	p8 := NewPoint(3, 4)
	fmt.Printf("Расстояние между %s и %s: %.2f\n", p7, p8, Distance(p7, p8))

	p9 := NewPoint(1.5, 2.7)
	p10 := NewPoint(4.2, 6.9)
	fmt.Printf("Расстояние между %s и %s: %.4f\n", p9, p10, Distance(p9, p10))
}
