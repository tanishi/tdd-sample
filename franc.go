package main

type Franc struct {
	*MoneyStruct
}

func NewFranc(amount int) *Franc {
	return &Franc{
		&MoneyStruct{
			amount: amount,
		},
	}
}

func (f *Franc) Times(multplier int) *Franc {
	return NewFranc(f.amount * multplier)
}

func (f *Franc) Equals(f2 *Franc) bool {
	return f.amount == f2.amount
}
