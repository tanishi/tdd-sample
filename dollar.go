package main

type Dollar struct {
	*MoneyStruct
}

func NewDollar(amount int) Money {
	return &Dollar{
		&MoneyStruct{
			amount: amount,
		},
	}
}

func (d *Dollar) Times(multplier int) Money {
	return NewDollar(d.amount * multplier)
}
