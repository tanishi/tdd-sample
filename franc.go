package main

type Franc struct {
	*MoneyStruct
}

func NewFranc(amount int, currency string) Money {
	return &Franc{
		NewMoney(amount, currency),
	}
}

func (f *Franc) Times(multplier int) Money {
	return NewFranc(f.amount*multplier, "CHF")
}
