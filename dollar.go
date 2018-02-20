package main

func NewDollar(amount int, currency string) *Money {
	return NewMoney(amount, currency)
}
