package main

type Franc struct {
	*MoneyStruct
}

func NewFranc(amount int, currency string) Money {
	return NewMoney(amount, currency)
}
