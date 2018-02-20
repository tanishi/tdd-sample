package main

func NewFranc(amount int, currency string) *Money {
	return NewMoney(amount, currency)
}
