package main

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Equals(obj interface{}) bool {
	money := obj.(*Money)
	return m.Amount() == money.Amount() && m.Currency() == money.Currency()
}

func (m *Money) Times(multplier int) *Money {
	return NewMoney(m.amount*multplier, m.Currency())
}
