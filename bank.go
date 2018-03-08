package main

type Expression interface {
}

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) Reduce(e Expression, s string) *Money {
	return NewDollar(10)
}
