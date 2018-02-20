package main

import (
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewMoney(5, "USD")

	expected := NewMoney(10, "USD")

	if !reflect.DeepEqual(five.Times(2), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(2))
	}

	expected = NewMoney(15, "USD")

	if !reflect.DeepEqual(five.Times(3), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(3))
	}
}

func TestEquality(t *testing.T) {
	if !NewMoney(5, "USD").Equals(NewMoney(5, "USD")) {
		t.Errorf("Equals error")
	}

	if NewMoney(5, "USD").Equals(NewMoney(6, "USD")) {
		t.Errorf("Equals error")
	}

	if !NewMoney(5, "CHF").Equals(NewMoney(5, "CHF")) {
		t.Errorf("Equals error")
	}

	if NewMoney(5, "CHF").Equals(NewMoney(6, "CHF")) {
		t.Errorf("Equals error")
	}

	if !NewMoney(5, "CHF").Equals(NewMoney(5, "CHF")) {
		t.Errorf("Equals error")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewMoney(5, "CHF")

	expected := NewMoney(10, "CHF")

	if !reflect.DeepEqual(five.Times(2), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(2))
	}

	expected = NewMoney(15, "CHF")

	if !reflect.DeepEqual(five.Times(3), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(3))
	}
}

func TestCurrency(t *testing.T) {
	if "USD" != NewMoney(1, "USD").Currency() {
		t.Errorf("expected: %v, but %v", "USD", NewMoney(1, "USD").Currency())
	}

	if "CHF" != NewMoney(1, "CHF").Currency() {
		t.Errorf("expected: %v, but %v", "CHF", NewMoney(1, "CHF").Currency())
	}
}
