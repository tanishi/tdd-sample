package main

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)

	expected := 10

	if product.Amount != expected {
		t.Errorf("expected: %v, but: %v", expected, product.Amount)
	}

	product = five.Times(3)
	expected = 15

	if product.Amount != expected {
		t.Errorf("expected: %v, but: %v", expected, product.Amount)
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).Equals(NewDollar(5)) {
		t.Errorf("Equals error")
	}

	if NewDollar(5).Equals(NewDollar(6)) {
		t.Errorf("Equals error")
	}
}
