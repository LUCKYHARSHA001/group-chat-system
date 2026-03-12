package main

import "fmt"

type BankAccount struct{
	Owner string
	Balance float64
}

func (a *BankAccount)Deposit(amount float64){
	a.Balance+=amount
}

func (a *BankAccount)WithDraw(amount float64) bool{
	if a.Balance<amount{
		return false
	}
	a.Balance-=amount
	return true
}

func main(){
	acc:=BankAccount{
		Owner: "Harsha",
		Balance: 100.0,
	}
	acc.Deposit(50.0)

	fmt.Printf("After Deposit:%.2f\n",acc.Balance)

	success:=acc.WithDraw(150.0)
	if success{
	fmt.Printf("After Withdraw:%.2f\n",acc.Balance)
	}else{
		fmt.Printf("Balance was low than expectd")
	}
}