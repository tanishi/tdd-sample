package main

type Pair struct {
	From string
	To   string
}

func NewPair(from, to string) *Pair {
	return &Pair{
		From: from,
		To:   to,
	}
}

func (p *Pair) Equals(obj interface{}) bool {
	o := obj.(Pair)

	return p.From == o.From && p.To == o.To
}

func (p *Pair) HashCode() int {
	return 0
}
