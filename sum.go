package main

type Sum struct {
	Augend Money
	Addend Money
}

func NewSum(augend, addend Money) *Sum {
	return &Sum{
		Augend: augend,
		Addend: addend,
	}
}

func (s *Sum) Reduce(to string) *Money {
	amount := s.Augend.Amount + s.Addend.Amount
	return NewMoney(amount, to)
}
