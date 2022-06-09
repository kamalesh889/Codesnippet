package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Context use in case of timeout")
	Ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go Handle(Ctx)
	select {
	case <-Ctx.Done():
		fmt.Println("deadline exceeded")
	}
	time.Sleep(2 * time.Second)

}

func Handle(ctx context.Context) {
	for { //For loop is demonstrating a long running process
		select {
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		default:
			fmt.Println("Work Process")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

//Sometimes have to send a response within a deadline or request will be failed

//So it can be used to control the system when deadline exceeded
