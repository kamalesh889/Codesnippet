package main

import "fmt"

type Connection struct {
	Address string
	port    string
}

func init() {
	C := Connection{
		Address: "werrr/adds",
		port:    "9090",
	}

	fmt.Println("Instance od connection created", C)
}

func main() {

}
