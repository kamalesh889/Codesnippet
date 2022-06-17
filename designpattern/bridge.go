package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float32
	getPerimeter() float32
	getColour()
	setColour(colour Colour)
}

type Square struct {
	side   float32
	colour Colour
}

func (s *Square) getArea() float32 {
	return s.side * s.side
}

func (s *Square) getPerimeter() float32 {
	return s.side * 4
}

func (s *Square) setColour(colour Colour) {
	s.colour = colour
}

func (s *Square) getColour() {
	fmt.Print("Square says: ")
	s.colour.printColour()
}

func NewSquare(side float32, colour Colour) *Square {
	s := &Square{side: side}
	s.setColour(colour)
	return s
}

type Circle struct {
	radius float32
	colour Colour
}

func (c *Circle) getArea() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) getPerimeter() float32 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) setColour(colour Colour) {
	c.colour = colour
}

func (c *Circle) getColour() {
	fmt.Print("Circle says: ")
	c.colour.printColour()
}

func NewCircle(radius float32, colour Colour) *Circle {
	c := &Circle{radius: radius}
	c.setColour(colour)
	return c
}

type Colour interface {
	printColour()
}

type RedColour struct{}

func (r *RedColour) printColour() {
	fmt.Println("I'm Red!")
}

type BlueColour struct{}

func (b *BlueColour) printColour() {
	fmt.Println("I'm Blue!")
}

func main() {
	r := &RedColour{}
	b := &BlueColour{}

	s := NewSquare(4.2, r)
	c := NewCircle(6.9, b)

	fmt.Println("Square's area is:", s.getArea())
	fmt.Println("Square's perimeter is:", s.getPerimeter())
	s.getColour()

	s.setColour(b)
	s.getColour()

	fmt.Println()

	fmt.Println("Circle's area is:", c.getArea())
	fmt.Println("Circle's perimeter is:", c.getPerimeter())
	c.getColour()

	c.setColour(r)
	c.getColour()
}

//It decouples abstraction (an object) from its implementation (the thing that the object does).
// This way, we can change what an object does as much as we want.
