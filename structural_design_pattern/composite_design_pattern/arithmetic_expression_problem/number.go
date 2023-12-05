package main

type Number struct {
	value int64
}

func NewNumber(value int64) Number {
	return Number{
		value: value,
	}
}

func (n Number) Evaluate() int64 {
	return n.value
}
