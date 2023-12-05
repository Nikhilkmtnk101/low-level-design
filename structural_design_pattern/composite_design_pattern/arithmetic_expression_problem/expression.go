package main

type Expression struct {
	operation      Operation
	leftOperation  ArithmeticExpression
	rightOperation ArithmeticExpression
}

func NewExpression(
	operation Operation, leftOperation, rightOperation ArithmeticExpression,
) *Expression {
	return &Expression{
		operation:      operation,
		leftOperation:  leftOperation,
		rightOperation: rightOperation,
	}
}

func (exp Expression) Evaluate() int64 {
	var value int64
	leftResult := exp.leftOperation.Evaluate()
	rightResult := exp.rightOperation.Evaluate()

	switch exp.operation {
	case Add:
		value = leftResult + rightResult
	case Subtract:
		value = leftResult - rightResult
	case Multiply:
		value = leftResult * rightResult
	case Divide:
		value = leftResult / rightResult
	}

	return value
}
