package main

import "fmt"

func main() {
	myWallet := Wallet{
		balance: 0,
	}
	chargePoint(&myWallet)
	myWallet.checkBalance()
	fmt.Println("Balance setelah mengecek saldo", myWallet.balance)
}
func chargePoint(wallet *Wallet) {
	wallet.balance = 100
}

func (wallet *Wallet) checkBalance() {
	fmt.Println("Mengecek saldo dikenakan biaya sebesar 10")
	fmt.Println("Saldo Total :", wallet.balance)
	wallet.balance -= 10
}

type Wallet struct {
	balance int
}
