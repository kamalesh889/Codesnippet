package main

import "fmt"

//About pipiline and how can we create it and use it

//Squaring of numbers by creating pipeline

func main() {
	fmt.Println("Pipeline start")
	Num := input(2, 3)
	op := square(Num)
	for i := range op {
		fmt.Println(i)
	}

}

func input(num ...int) <-chan int {
	x := make(chan int)

	go func() {
		for i := range num {
			x <- num[i]
		}
		close(x)
		//Always close the channel that you write into otherwise the listener listens for infinity resulting in panic

	}()
	return x
}

func square(nos <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for i := range nos {
			output <- i * i
		}
		close(output)

	}()
	return output
}

//Informally, a pipeline is a series of stages connected by channels, where each stage is a group of goroutines running the same function. In each stage, the goroutines
// receive values from upstream via inbound channels
// perform some function on that data, usually producing new values
// send values downstream via outbound channels

// -------------------------------------------------------------------
// -------------------------------------------------------------------

// data pipelines are a sequence of steps that are executed using the output of a previous step as the input for the next one. That sound like a really good usage for goroutines and channels. We can create each step as a goroutine, and the communications between them as channels
