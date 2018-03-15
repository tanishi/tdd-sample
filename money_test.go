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
	}

	five := NewDollar(5)
	bank := NewBank()

	for _, c := range cases {
		sum := five.Plus(c.in)
		got := bank.Reduce(sum, "USD")

		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("want: %v, but: %v", c.want, got)
		}
	}
}

func TestEquality(t *testing.T) {
	cases := []struct {
		lfs  *Money
		rhs  *Money
		want bool
	}{
		{NewDollar(5), NewDollar(5), true},
		{NewDollar(5), NewDollar(6), false},
		{NewFranc(5), NewFranc(5), true},
		{NewFranc(5), NewFranc(6), false},
		{NewDollar(5), NewFranc(5), false},
	}

	for _, c := range cases {
		if c.lfs.Equals(c.rhs) != c.want {
			t.Errorf("Equals error")
		}
	}
}

func TestCurrency(t *testing.T) {
	cases := []struct {
		in   *Money
		want string
	}{
		{NewDollar(1), "USD"},
		{NewFranc(1), "CHF"},
	}

	for _, c := range cases {
		got := c.in.Currency

		if c.want != got {
			t.Errorf("want: %v, but %v", c.want, got)
		}
	}
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum := result.(*Sum)

	if !reflect.DeepEqual(*five, sum.Augend) {
		t.Errorf("want: %v, but %v", *five, sum.Augend)
	}

	if !reflect.DeepEqual(*five, sum.Addend) {
		t.Errorf("want: %v, but %v", *five, sum.Addend)
	}
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()

	want := NewDollar(1)
	got := bank.Reduce(NewDollar(1), "USD")

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)

	want := NewDollar(1)
	got := bank.Reduce(NewFranc(2), "USD")

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestIdentityRate(t *testing.T) {
	want := 1
	got := NewBank().Rate("USD", "USD")

	if want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestMixAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)

	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)

	want := NewDollar(10)
	got := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
