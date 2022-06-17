package main

import "fmt"

//Start
type pizza interface {
	getPrice() int
}

type veg struct{}

func (v *veg) getPrice() int {
	return 100
}

type Nonveg struct{}

func (n *Nonveg) getPrice() int {
	return 200
}

//Upto this old functions

//Now we have new funcs.. So this pattern will help us to implement those
//Without altering the objects

type Tomatotopping struct {
	piz pizza
}

func (t *Tomatotopping) getPrice() int {
	return t.piz.getPrice() + 50
}

type Cheesetopping struct {
	pi pizza
}

func (c *Cheesetopping) getPrice() int {
	return c.pi.getPrice() + 70
}

func main() {
	vegpizza := &veg{}
	// Price := vegpizza.getPrice()
	// fmt.Println(Price)
	Tomatopiz := &Tomatotopping{
		piz: vegpizza,
	}
	Price := Tomatopiz.getPrice()
	fmt.Println("Tomoato with vegpizza", Price)

	Nonvpizza := &Nonveg{}
	Cheesepiz := &Cheesetopping{
		pi: Nonvpizza,
	}
	Price = Cheesepiz.getPrice()
	fmt.Println("Chesese with nonveg", Price)
}

//The Decorator design pattern allows you to decorate an already existing type
//with more functional features without actually touching it.
