package main

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		amount: amount,
	}
}

func (d *Dollar) Times(multplier int) *Dollar {
	return NewDollar(d.amount * multplier)
}

func (d *Dollar) Equals(d2 *Dollar) bool {
	return d.amount == d2.amount
}
