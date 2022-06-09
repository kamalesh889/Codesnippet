package main

import (
	"fmt"
)

type Account struct { //Suppose a complex struct
	Acc_no  int
	Name    string
	Branch  string
	Balance float64
}

//Account Constructor
func Create_account() Account {
	return Account{}
}

//Builders
func (A Account) Set_Accno(number int) Account {
	A.Acc_no = number
	return A
}

func (A Account) Set_Name(name string) Account {
	A.Name = name
	return A
}

func (A Account) Set_Branch(branch string) Account {
	A.Branch = branch
	return A
}
func (A Account) Set_Balance(balance float64) Account {
	A.Balance = balance
	return A
}

func main() {
	A := Create_account().Set_Accno(12344).Set_Name("Kamalesh").Set_Balance(100).Set_Branch("Puri")
	fmt.Printf("%+v", A)
}

//Builder pattern is to create a complex object piece-by-piece
//It aslo encapsulate the construction logic for an object
