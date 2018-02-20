package main

type Dollar struct {
	*MoneyStruct
}

func NewDollar(amount int, currency string) Money {
	return &Dollar{
		NewMoney(amount, currency),
	}
}

func (d *Dollar) Times(multplier int) Money {
	return NewDollar(d.amount*multplier, "USD")
}
