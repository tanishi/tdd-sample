package main

type Expression interface {
	Reduce(Bank, string) *Money
}
