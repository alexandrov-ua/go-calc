package parser

import (
	"math"
)

type SyntaxNode interface {
	Exec() float64
}

// ======LiteralNode
type LiteralNode struct {
	Value float64
}

func (n LiteralNode) Exec() float64 {
	return n.Value
}

// ======PlusBinaryOperationNode
type PlusBinaryOperationNode struct {
	Left, Right SyntaxNode
}

func (n PlusBinaryOperationNode) Exec() float64 {
	return n.Left.Exec() + n.Right.Exec()
}

// ======MinusBinaryOperationNode
type MinusBinaryOperationNode struct {
	Left, Right SyntaxNode
}

func (n MinusBinaryOperationNode) Exec() float64 {
	return n.Left.Exec() - n.Right.Exec()
}

// ======PlusUnaryOperationNode
type PlusUnaryOperationNode struct {
	Node SyntaxNode
}

func (n PlusUnaryOperationNode) Exec() float64 {
	return n.Node.Exec()
}

// ======MinusUnaryOperationNode
type MinusUnaryOperationNode struct {
	Node SyntaxNode
}

func (n MinusUnaryOperationNode) Exec() float64 {
	return n.Node.Exec() * -1
}

// ======MultiplyBinaryOperationNode
type MultiplyBinaryOperationNode struct {
	Left, Right SyntaxNode
}

func (n MultiplyBinaryOperationNode) Exec() float64 {
	return n.Left.Exec() * n.Right.Exec()
}

// ======DevideBinaryOperationNode
type DevideBinaryOperationNode struct {
	Left, Right SyntaxNode
}

func (n DevideBinaryOperationNode) Exec() float64 {
	return n.Left.Exec() / n.Right.Exec()
}

// ======PowerBinaryOperationNode
type PowerBinaryOperationNode struct {
	Left, Right SyntaxNode
}

func (n PowerBinaryOperationNode) Exec() float64 {
	return math.Pow(n.Left.Exec(), n.Right.Exec())
}

// ======ParenthesesNode
type ParenthesesNode struct {
	Node SyntaxNode
}

func (n ParenthesesNode) Exec() float64 {
	return n.Node.Exec()
}
