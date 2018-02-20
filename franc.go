package main

type Franc struct {
	*MoneyStruct
}

func NewFranc(amount int) Money {
	return &Franc{
		&MoneyStruct{
			amount: amount,
		},
	}
}

func (f *Franc) Times(multplier int) Money {
	return NewFranc(f.amount * multplier)
}

func (f *Franc) Currency() string {
	return "CHF"
}
