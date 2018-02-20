package main

type Money interface {
	Equals(obj interface{}) bool
	GetAmount() int
	Times(int) Money
	Currency() string
}

type MoneyStruct struct {
	amount   int
	currency string
}

func (m *MoneyStruct) GetAmount() int {
	return m.amount
}

func (m *MoneyStruct) Currency() string {
	return m.currency
}

func (m *MoneyStruct) Equals(obj interface{}) bool {
	money := obj.(Money)
	return m.GetAmount() == money.GetAmount()
}
