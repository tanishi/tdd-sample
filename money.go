package main

type Money struct {
	Amount   int
	Currency string
}

func NewMoney(amount int, currency string) *Money {
	return &Money{
		Amount:   amount,
		Currency: currency,
	}
}

func (m *Money) Equals(obj interface{}) bool {
	money := obj.(*Money)
	return m.Amount == money.Amount && m.Currency == money.Currency
}

func (m *Money) Times(multplier int) *Money {
	return NewMoney(m.Amount*multplier, m.Currency)
}
