package main

import (
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	cases := []struct {
		in   int
		want *Money
	}{
		{2, NewDollar(10)},
		{3, NewDollar(15)},
	}

	five := NewDollar(5)

	for _, c := range cases {
		got := five.Times(c.in)

		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("want: %v, but: %v", c.want, got)
		}
	}
}

func TestAddition(t *testing.T) {
	cases := []struct {
		in   *Money
		want *Money
	}{
		{NewDollar(5), NewDollar(10)},
		{NewDollar(7), NewDollar(12)},
	}

	five := NewDollar(5)

	for _, c := range cases {
		got := five.Plus(c.in)

		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("want: %v, but: %v", c.want, got)
		}
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).Equals(NewDollar(5)) {
		t.Errorf("Equals error")
	}

	if NewDollar(5).Equals(NewDollar(6)) {
		t.Errorf("Equals error")
	}

	if !NewFranc(5).Equals(NewFranc(5)) {
		t.Errorf("Equals error")
	}

	if NewFranc(5).Equals(NewFranc(6)) {
		t.Errorf("Equals error")
	}

	if !NewFranc(5).Equals(NewFranc(5)) {
		t.Errorf("Equals error")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)

	want := NewFranc(10)

	if !reflect.DeepEqual(five.Times(2), want) {
		t.Errorf("want: %v, but: %v", want, five.Times(2))
	}

	want = NewFranc(15)

	if !reflect.DeepEqual(five.Times(3), want) {
		t.Errorf("want: %v, but: %v", want, five.Times(3))
	}
}

func TestCurrency(t *testing.T) {
	if "USD" != NewDollar(1).Currency {
		t.Errorf("want: %v, but %v", "USD", NewDollar(1).Currency)
	}

	if "CHF" != NewFranc(1).Currency {
		t.Errorf("want: %v, but %v", "CHF", NewFranc(1).Currency)
	}
}
