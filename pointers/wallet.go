package wallet

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	//fmt.Printf("address in method is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
