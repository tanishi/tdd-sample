package main

type Dollar struct {
	*MoneyStruct
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		&MoneyStruct{
			amount: amount,
		},
	}
}

func (d *Dollar) Times(multplier int) *Dollar {
	return NewDollar(d.amount * multplier)
}
