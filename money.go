package main

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		Amount: amount,
	}
}

func (d *Dollar) Times(multplier int) *Dollar {
	return NewDollar(d.Amount * multplier)
}
