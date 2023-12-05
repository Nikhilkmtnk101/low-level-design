package main

import "fmt"

func clientCode(expression ArithmeticExpression) int64 {
	return expression.Evaluate()
}

func main() {
	exp := NewExpression(
		Multiply,
		NewExpression(Add, NewNumber(3), NewNumber(2)),
		NewNumber(5))
	fmt.Println("result of expression (3 + 2) * 5 =", clientCode(exp))
}
