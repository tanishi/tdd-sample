package main

import "testing"

func TestHelloWorld(t *testing.T) {
	five := NewDollar(5)
	five.Times(2)

	expected := 10

	if five.Amount != expected {
		t.Errorf("expected: %v, but: %v", expected, five.Amount)
	}
}
