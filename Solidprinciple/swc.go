package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	Chan1 := Arrange(wg)
	Chan2 := Sum(wg, Chan1)
	Print(wg, Chan2)
	wg.Wait()
}

func Arrange(wg *sync.WaitGroup) chan int { //It generate numbers
	defer wg.Done()
	arr := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			arr <- i
		}
		close(arr)
	}()
	return arr
}

func Sum(wg *sync.WaitGroup, cha chan int) chan int { //It add all the numbers
	defer wg.Done()
	sum := make(chan int)
	op := 0
	go func() {
		for i := range cha {
			op = op + i
		}
		sum <- op
	}()
	return sum
}

func Print(wg *sync.WaitGroup, Sum chan int) { //it Prints the result
	defer wg.Done()
	x := <-Sum
	fmt.Println("Summation is ---", x)
}

//Maintaining individuallity that is each func is performing only one operation

//Swc :=== Srp implemented with concurrency
