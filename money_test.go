package main

import (
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5, "USD")

	expected := NewDollar(10, "USD")

	if !reflect.DeepEqual(five.Times(2), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(2))
	}

	expected = NewDollar(15, "USD")

	if !reflect.DeepEqual(five.Times(3), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(3))
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5, "USD").Equals(NewDollar(5, "USD")) {
		t.Errorf("Equals error")
	}

	if NewDollar(5, "USD").Equals(NewDollar(6, "USD")) {
		t.Errorf("Equals error")
	}

	if !NewFranc(5, "CHF").Equals(NewFranc(5, "CHF")) {
		t.Errorf("Equals error")
	}

	if NewFranc(5, "CHF").Equals(NewFranc(6, "CHF")) {
		t.Errorf("Equals error")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5, "CHF")

	expected := NewFranc(10, "CHF")

	if !reflect.DeepEqual(five.Times(2), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(2))
	}

	expected = NewFranc(15, "CHF")

	if !reflect.DeepEqual(five.Times(3), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(3))
	}
}

func TestCurrency(t *testing.T) {
	if "USD" != NewDollar(1, "USD").Currency() {
		t.Errorf("expected: %v, but %v", "USD", NewDollar(1, "USD").Currency())
	}

	if "CHF" != NewFranc(1, "CHF").Currency() {
		t.Errorf("expected: %v, but %v", "CHF", NewFranc(1, "CHF").Currency())
	}
}
