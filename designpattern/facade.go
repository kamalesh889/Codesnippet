package main

import (
	"errors"
	"fmt"
)

type account struct {
	name string
}

func New_account(acc_name string) *account {
	return &account{
		name: acc_name,
	}
}

func (a *account) Check_account(acc_name string) error {
	if a.name != acc_name {
		return errors.New("account name is incorrect")
	}
	fmt.Println("Account name verified")
	return nil
}

type securitycode struct {
	code int
}

func New_securitycode(codegiven int) *securitycode {
	return &securitycode{
		code: codegiven,
	}
}
func (s *securitycode) Check_code(codegiven int) error {
	if s.code != codegiven {
		return errors.New("Code mismatched")
	}
	fmt.Println("Code matched")
	return nil
}

type wallet struct {
	balance int
}

func New_wallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) Credit_wallet(amount int) {
	w.balance += amount
	fmt.Println("Amount added into account", "Updated balance-->", w.balance)
	return
}

func (w *wallet) Debit_wallet(amount int) error {
	if w.balance < amount {
		return errors.New("Sufficient balance not available")
	}
	w.balance -= amount
	fmt.Println("Amount deducted from account", "Remaining balace-->", w.balance)
	return nil
}

type Walletfacade struct {
	account      *account
	securitycode *securitycode
	wallet       *wallet
}

func newWalletFacade(Name string, code int) *Walletfacade {
	fmt.Println("Starting create account")
	Wf := &Walletfacade{
		account:      New_account(Name),
		securitycode: New_securitycode(code),
		wallet:       New_wallet(),
	}
	fmt.Println("Account Created")
	return Wf
}

func (wf *Walletfacade) Addmoney(Name string, code int, Amount int) error {
	fmt.Println("Starting to add money into the wallet")
	err := wf.account.Check_account(Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = wf.securitycode.Check_code(code)
	if err != nil {
		fmt.Println(err)
		return err
	}
	wf.wallet.Credit_wallet(Amount)
	return nil
}

func (wf *Walletfacade) Deductmoney(Name string, code int, Amount int) error {
	fmt.Println("Starting to deduct money from the wallet")
	err := wf.account.Check_account(Name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = wf.securitycode.Check_code(code)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = wf.wallet.Debit_wallet(Amount)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	Newwf := newWalletFacade("Kamalesh", 123)
	err := Newwf.Addmoney("Kamalesh", 123, 1000)
	if err != nil {
		return
	}
	err = Newwf.Deductmoney("Kamalesh", 123, 700)
	if err != nil {
		fmt.Println(err)
	}
}

//https://golangbyexample.com/facade-design-pattern-in-golang/

//his design pattern is meant to hide the complexities of the underlying
//system and provide a simple interface to the client
