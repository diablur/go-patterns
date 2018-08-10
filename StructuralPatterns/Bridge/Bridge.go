package main

import (
	"fmt"
)

type Drawer interface {
	DrawCircle(x, y, radius float32)
}

type OpenGL struct{}

func (gl *OpenGL) DrawCircle(x, y, radius float32) {
	fmt.Printf("OpenGL.circle at %f:%f radius %f\n", x, y, radius)
}

type Direct2D struct{}

func (d2d *Direct2D) DrawCircle(x, y, radius float32) {
	fmt.Printf("Direct2D.circle at %f:%f radius %f\n", x, y, radius)
}

type Shape interface {
	Draw()
}

type Circle struct {
	x, y, radius   float32
	drawingContext Drawer
}

func (s *Circle) Draw() {
	s.drawingContext.DrawCircle(s.x, s.y, s.radius)
}

func main() {
	shapes := []Shape{
		&Circle{1, 2, 3, new(OpenGL)},
		&Circle{5, 7, 11, new(Direct2D)},
	}
	for _, shape := range shapes {
		shape.Draw()
	}
}
