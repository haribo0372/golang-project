package main

import (
	"fmt"
	m "math"
)

type Point struct {
	X, Y float64
}

type Section struct {
	Borders [2]Point
}

func (s Section) Length() float64 {
	// sqrt (x2 — x1)^2 + (y2 — y1)^2
	return m.Sqrt(m.Pow(s.Borders[0].X-s.Borders[1].X, 2) + m.Pow(s.Borders[0].Y-s.Borders[1].Y, 2))
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	A, B, C Point
}

type Circle struct {
	Center Point
	Radius float64
}

func (t Triangle) Area() float64 {
	return m.Abs((t.A.X*(t.B.Y-t.C.Y) + t.B.X*(t.C.Y-t.A.Y) + t.C.X*(t.A.Y-t.B.Y)) / 2)
}

func (c Circle) Area() float64 {
	return m.Pi * m.Pow(c.Radius, 2)
}

func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}
