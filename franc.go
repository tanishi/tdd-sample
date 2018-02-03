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
