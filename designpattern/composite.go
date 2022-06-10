package main

import "fmt"

//-----------------------------------
type ath struct{}

func (a *ath) train() {
	fmt.Println("Training")
}

type compswimmer struct {
	athlete ath
	myswim  func()
}

func swim() {
	fmt.Println("Swimming")
}

//-----------------------------------

type animal struct{}

func (a *animal) eat() {
	fmt.Println("Eating")
}

type fish struct {
	animal //emebed object within onject
	myswim func()
}

//-----------------------------------

type swimmer interface {
	swim()
}
type atheletic interface {
	train()
}

type swimimple struct{}

func (s *swimimple) swim() {
	fmt.Println("swimming")
}

type composite struct {
	swimmer
	atheletic
}

func main() {
	//One
	// c := compswimmer{
	// 	myswim: swim, //If not pass athlete,it will create one due to zero initialization
	// }
	// c.athlete.train()

	//two
	// f := fish{
	// 	myswim: swim,
	// }
	// f.eat()

	//three
	c := composite{
		&swimimple{},
		&ath{},
	}
	c.swim()
	c.train()

}
