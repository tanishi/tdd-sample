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

func (s *Sum) Reduce(b Bank, to string) *Money {
	amount := s.Augend.Reduce(b, to).Amount + s.Addend.Reduce(b, to).Amount
	return NewMoney(amount, to)
}
