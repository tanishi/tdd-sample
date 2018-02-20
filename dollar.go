package main

type Dollar struct {
	*MoneyStruct
}

func NewDollar(amount int, currency string) Money {
	return NewMoney(amount, currency)
}
