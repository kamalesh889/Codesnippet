package main

import "fmt"

func main() {
	square := Nos{length: 5}
	print(square.area())

}

type Nos struct {
	length int
}

func (c Nos) area() int {
	return c.length * c.length
}

func print(result int) {
	fmt.Println("result is ", result)
}

//Area and print different functions executig different tasks
