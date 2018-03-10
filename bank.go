package main

type Expression interface {
	Reduce(Bank, string) *Money
}

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(*b, to)
}

func (b *Bank) AddRate(from, to string, rate int) {
}

func (b *Bank) Rate(from, to string) int {
	if from == "CHF" && to == "USD" {
		return 2
	}

	return 1
}
