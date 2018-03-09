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
	money, ok := obj.(*Money)

	if !ok {
		return false
	}

	return m.Amount == money.Amount && m.Currency == money.Currency
}

func (m *Money) Plus(addend *Money) Expression {
	return NewSum(*m, *addend)
}

func (m *Money) Times(multplier int) *Money {
	return NewMoney(m.Amount*multplier, m.Currency)
}

func (m *Money) Reduce(to string) *Money {
	return m
}

func NewDollar(amount int) *Money {
	return NewMoney(amount, "USD")
}

func NewFranc(amount int) *Money {
	return NewMoney(amount, "CHF")
}
