package main

import (
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)

	expected := NewDollar(10)

	if !reflect.DeepEqual(five.Times(2), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(2))
	}

	expected = NewDollar(15)

	if !reflect.DeepEqual(five.Times(3), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(3))
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

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)

	expected := NewFranc(10)

	if !reflect.DeepEqual(five.Times(2), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(2))
	}

	expected = NewFranc(15)

	if !reflect.DeepEqual(five.Times(3), expected) {
		t.Errorf("expected: %v, but: %v", expected, five.Times(3))
	}
}
