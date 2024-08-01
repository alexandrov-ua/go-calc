package parser

import (
	"testing"
)

type UnarySyntaxNode interface {
	GetNode() SyntaxNode
}

type BynarySyntaxNode interface {
	GetLeftAndRight() (SyntaxNode, SyntaxNode)
}

func (n DevideBinaryOperationNode) GetLeftAndRight() (SyntaxNode, SyntaxNode) {
	return n.Left, n.Right
}

func (n MultiplyBinaryOperationNode) GetLeftAndRight() (SyntaxNode, SyntaxNode) {
	return n.Left, n.Right
}

func (n PlusBinaryOperationNode) GetLeftAndRight() (SyntaxNode, SyntaxNode) {
	return n.Left, n.Right
}

func (n MinusBinaryOperationNode) GetLeftAndRight() (SyntaxNode, SyntaxNode) {
	return n.Left, n.Right
}

func (n PowerBinaryOperationNode) GetLeftAndRight() (SyntaxNode, SyntaxNode) {
	return n.Left, n.Right
}

func (n PlusUnaryOperationNode) GetNode() SyntaxNode {
	return n.Node
}

func (n MinusUnaryOperationNode) GetNode() SyntaxNode {
	return n.Node
}

func (n ParenthesesNode) GetNode() SyntaxNode {
	return n.Node
}

func MatchNode[T SyntaxNode](t *testing.T, s SyntaxNode) {
	if _, ok := s.(T); !ok {
		t.Errorf("Expecting %T but found: %T", LiteralNode{}, s)
	}
}

func MatchValueNode(t *testing.T, s SyntaxNode, val float64) {
	if v, ok := s.(LiteralNode); !ok {
		t.Errorf("Expecting %T but found: %T", LiteralNode{}, s)
	} else if v.Value != val {
		t.Errorf("Expecting %v but found: %v", val, v.Value)
	}
}

func MatchUnaryNode[T UnarySyntaxNode](t *testing.T, s SyntaxNode) SyntaxNode {
	if n, ok := s.(T); !ok {
		t.Errorf("Expecting %T but found: %T", LiteralNode{}, s)
		return nil
	} else {
		return n.GetNode()
	}
}

func MatchBinaryNode[T BynarySyntaxNode](t *testing.T, s SyntaxNode) (SyntaxNode, SyntaxNode) {
	if n, ok := s.(T); !ok {
		t.Errorf("Expecting %T but found: %T", LiteralNode{}, s)
		return nil, nil
	} else {
		return n.GetLeftAndRight()
	}
}

func Test_Positive(t *testing.T) {
	res, _ := Parse("123")
	MatchNode[LiteralNode](t, res)
}

func Test_Unary_Positive(t *testing.T) {
	res, _ := Parse("-123")
	res = MatchUnaryNode[MinusUnaryOperationNode](t, res)
	MatchNode[LiteralNode](t, res)
}

func Test_Binary_Positive(t *testing.T) {
	res, _ := Parse("1+2")
	left, right := MatchBinaryNode[PlusBinaryOperationNode](t, res)
	MatchNode[LiteralNode](t, left)
	MatchNode[LiteralNode](t, right)
}

func Test_Binary_Precedence(t *testing.T) {
	res, _ := Parse("2*3+4")
	left, right := MatchBinaryNode[PlusBinaryOperationNode](t, res)
	MatchValueNode(t, right, 4)
	left, right = MatchBinaryNode[MultiplyBinaryOperationNode](t, left)
	MatchValueNode(t, left, 2)
	MatchValueNode(t, right, 3)
}

func Test_Binary_Precedence_2(t *testing.T) {
	res, _ := Parse("2^5*3+4")
	left, right := MatchBinaryNode[PlusBinaryOperationNode](t, res)
	MatchValueNode(t, right, 4)
	left, right = MatchBinaryNode[MultiplyBinaryOperationNode](t, left)
	MatchValueNode(t, right, 3)
	left, right = MatchBinaryNode[PowerBinaryOperationNode](t, left)
	MatchValueNode(t, left, 2)
	MatchValueNode(t, right, 5)
}

func Test_Binary_Precedence_Parentheses(t *testing.T) {
	res, _ := Parse("2*(3+4)")
	left, right := MatchBinaryNode[MultiplyBinaryOperationNode](t, res)
	MatchValueNode(t, left, 2)
	tmp := MatchUnaryNode[ParenthesesNode](t, right)
	left, right = MatchBinaryNode[PlusBinaryOperationNode](t, tmp)
	MatchValueNode(t, left, 3)
	MatchValueNode(t, right, 4)
}
