package main

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{
		amount: amount,
	}
}

func (f *Franc) Times(multplier int) *Franc {
	return NewFranc(f.amount * multplier)
}

func (f *Franc) Equals(f2 *Franc) bool {
	return f.amount == f2.amount
}
