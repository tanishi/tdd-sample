package main

type Money interface {
	Equals(obj interface{}) bool
	Times(int) Money
	Amount() int
	Currency() string
}

type MoneyStruct struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) *MoneyStruct {
	return &MoneyStruct{
		amount:   amount,
		currency: currency,
	}
}

func (m *MoneyStruct) Amount() int {
	return m.amount
}

func (m *MoneyStruct) Currency() string {
	return m.currency
}

func (m *MoneyStruct) Equals(obj interface{}) bool {
	money := obj.(Money)
	return m.Amount() == money.Amount() && m.Currency() == money.Currency()
}

func (m *MoneyStruct) Times(multplier int) Money {
	return NewMoney(m.amount*multplier, m.Currency())
}
