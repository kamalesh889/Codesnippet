package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Context usecase")
	http.HandleFunc("/Handle", Handel)
	http.ListenAndServe(":8080", nil)
}

func Handel(res http.ResponseWriter, req *http.Request) {
	Con := req.Context()
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(res, "Work processed")
	case <-Con.Done():
		fmt.Println("error is", Con.Err().Error())
		return
	}
}

// If we cancel the reqest with in first 5 seconds , Context is used to handle this cancellation of request.
