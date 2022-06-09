package main

type Account struct {
	name    string
	no      int
	balance float64
}

type Transaction struct {
	Rec_acc Account
	sender  Account
	Id      int64
	Amount  int
}

func (A Account) Update_name() {

}

func (B Account) Credit(a int) {

}

func (C Account) Debit(x int) {

}

func (T Transaction) Make_transfer(a int) {
	T.sender.Debit(a)
	T.Rec_acc.Credit(a)
	T.Amount = a

}

func main() {
	A1 := Account{"Kamlesh", 99, 8080}
	A2 := Account{"Sahil", 90, 808000}
	T1 := Transaction{A1, A2, 999999, 0}
	T1.Make_transfer(20)
}
