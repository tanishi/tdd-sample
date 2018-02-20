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

func (m *MoneyStruct) Amount() int {
	return m.amount
}

func (m *MoneyStruct) Currency() string {
	return m.currency
}

func (m *MoneyStruct) Equals(obj interface{}) bool {
	money := obj.(Money)
	return m.Amount() == money.Amount()
}
