package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Scenario creating person object for 2 employees of the same company
type Person struct {
	Name        string
	Age         int
	WorkAddress *WorkAddress
	HomeAddress *HomeAddress
}

type WorkAddress struct {
	CompanyName, City, Country string
}

type HomeAddress struct {
	HouseNumber, City, Country string
}

func (p *Person) DeepCopy() *Person {
	pe := &Person{}
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	enc.Encode(p)
	dec := gob.NewDecoder(&buf)
	dec.Decode(pe)
	return pe
}

func main() {
	// Suppose we have a person object
	person1 := &Person{
		Name: "John",
		Age:  30,
		WorkAddress: &WorkAddress{
			CompanyName: "abc",
			City:        "London",
			Country:     "UK",
		},
		HomeAddress: &HomeAddress{
			HouseNumber: "1A",
			City:        "London",
			Country:     "UK",
		},
	}

	person2 := person1.DeepCopy()

	// Modifying name and house address
	person2.Name = "Surya"
	person2.HomeAddress.HouseNumber = "134C"
	person2.HomeAddress.City = "Bangalore"
	person2.HomeAddress.Country = "India"

	fmt.Println(person1, person1.WorkAddress, person1.HomeAddress)
	fmt.Println(person2, person2.WorkAddress, person2.HomeAddress)
}

////https://devcharmander.medium.com/design-pattern-in-golang-prototype-e864522e4eeb
