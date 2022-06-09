package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Ways for implementing concurrency ")
	// Case1()
	// Case2()
	//Case3()
}

//Case 1
// Use of time.sleep (It will sleep for time interval)
func Case1() {
	fmt.Println("Use of time.sleep")
	go Start()
	go End()
	time.Sleep(1 * time.Second)
}
func Start() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func End() {
	fmt.Println("End")
}

//Case 2
//Use of Wait group  (Provides sync mechanisim to stop the execution of program untill all the go routines released in the group)
func Case2() {
	fmt.Println("Use of Wait group")
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go Start2(wg)
	go End2(wg)
	wg.Wait()

	// for i := 0; i < 3; i++ {   // Use of multiple go rutines can be achived by this way
	// 	wg.Add(2)
	// 	go Start2(wg)
	// 	go End2(wg)
	// }
	// wg.Wait()
}

func Start2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}

func End2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("End")
}

//Case 3
//Use of channels (Used to establish communication betweem channels)

var Channel = make(chan bool)

func Case3() {
	go Start3()
	<-Channel
	go End3()
	<-Channel

}

func Start3() {
	for i := 0; i < 60; i++ {
		fmt.Println(i)
	}
	Channel <- true
}

func End3() {
	fmt.Println("End")
	Channel <- true
}
