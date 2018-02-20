package main

type Franc struct {
	*MoneyStruct
}

func NewFranc(amount int, currency string) Money {
	return &Franc{
		&MoneyStruct{
			amount:   amount,
			currency: currency,
		},
	}
}

func (f *Franc) Times(multplier int) Money {
	return NewFranc(f.amount*multplier, "CHF")
}
