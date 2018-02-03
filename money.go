package main

type Money interface {
	Equals(obj interface{}) bool
	GetAmount() int
}

type MoneyStruct struct {
	amount int
}

func (m *MoneyStruct) GetAmount() int {
	return m.amount
}

func (m *MoneyStruct) Equals(obj interface{}) bool {
	money := obj.(Money)
	return m.GetAmount() == money.GetAmount()
}
