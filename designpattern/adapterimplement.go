package main

import "fmt"

//---------------------------------------
//Old interface

type ChangeOld interface {
	Change_name(string)
}

type Chold struct {
	name string
}

func (c *Chold) Change_name(s string) {
	fmt.Println("name changed", s)

}

//---------------------------------------
//New interface
type Changenew interface {
	Change_address() string
}

//---------------------------------------
//Adapter
type Adapter struct {
	Oldone ChangeOld
	text   string
}

func (a *Adapter) Change_addres() {
	if a.Oldone != nil {
		a.Oldone.Change_name(a.text)
	} else {
		fmt.Println("address chnaged", a.text)
	}

}

//---------------------------------------
//Account Creation
type Account struct {
	name    string
	number  int
	pin     int
	address string
}

func (A *Account) Account_details() {
	fmt.Println("Account is created ", A)
}

func main() {
	A := &Account{
		name:    "kamli",
		number:  9090889,
		pin:     752001,
		address: "BBsr",
	}
	A.Account_details()

	B := &Adapter{
		Oldone: &Chold{},
		text:   "Ram",
	}
	B.Change_addres()

	C := &Adapter{
		Oldone: nil,
		text:   "Puri",
	}
	C.Change_addres()

}
