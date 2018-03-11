package main

type Expression interface {
	Reduce(Bank, string) *Money
}

type Bank struct {
	rate map[Pair]int
}

func NewBank() *Bank {
	return &Bank{
		rate: make(map[Pair]int),
	}
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(*b, to)
}

func (b *Bank) AddRate(from, to string, rate int) {
	b.rate[*NewPair(from, to)] = rate
}

func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	return b.rate[*NewPair(from, to)]
}
