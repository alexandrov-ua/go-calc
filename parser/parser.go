package parser

import (
	"calc/lexer"
	"fmt"
	"strconv"
)

func Parse(s string) (SyntaxNode, []ParserDiagnostic) {
	ctx := ParserContext{Lexer: *lexer.Create(s), Diagnostics: make([]ParserDiagnostic, 0)}
	return ctx.Parse()
}

type ParserContext struct {
	Lexer       lexer.Lexer
	Diagnostics []ParserDiagnostic
}

func (p *ParserContext) Parse() (SyntaxNode, []ParserDiagnostic) {
	root := p.ParseBinaryExpression(0)
	_, _ = p.matchToken(lexer.EndOfFile)
	return root, p.Diagnostics
}

func (p *ParserContext) ParseBinaryExpression(precedence int) SyntaxNode {
	left := p.ParseUnaryOrLiteral()
	for p.Lexer.Current.Type != lexer.EndOfFile && p.Lexer.Current.Type != lexer.CloseParentheses {
		if pr := GetOperatorPrecedence(p.Lexer.Current); pr < precedence {
			break
		} else {
			op, ok := p.matchToken(lexer.Operator)
			if !ok {
				break
			}
			right := p.ParseBinaryExpression(pr)
			left = CreateBinaryOperation(op, left, right)
		}
	}
	return left
}

func (p *ParserContext) ParseUnaryOrLiteral() SyntaxNode {
	currToken := p.Lexer.Current
	switch {
	case currToken.Type == lexer.OpenParentheses:
		return p.ParseParentheses()
	case currToken.StrVal == "+" || currToken.StrVal == "-":
		return CreateUnaryOperation(p.Lexer.GetAndMooveNext(), p.ParseUnaryOrLiteral())
	default:
		return p.ParseLiteral()
	}
}

func (p *ParserContext) ParseParentheses() SyntaxNode {
	_, _ = p.matchToken(lexer.OpenParentheses)
	node := p.ParseBinaryExpression(0)
	_, _ = p.matchToken(lexer.CloseParentheses)
	return ParenthesesNode{node}
}

func (p *ParserContext) matchToken(tt lexer.TokenType) (lexer.Token, bool) {
	if p.Lexer.Current.Type != tt {
		p.Diagnostics = append(p.Diagnostics, ParserDiagnostic{fmt.Sprintf("Expected %s but found %v", tt, p.Lexer.Current.Type), p.Lexer.Current.Span})
		p.Lexer.MoveNext()
		return lexer.Token{Type: tt, Span: lexer.Span{Start: 0, End: 0}, StrVal: getDefoultStrValForError(tt)}, false
	}
	defer p.Lexer.MoveNext()
	return p.Lexer.Current, true
}

func getDefoultStrValForError(t lexer.TokenType) string {
	switch t {
	case lexer.Identifier, lexer.Literal:
		return "0"
	case lexer.Operator:
		return "+"
	default:
		return ""
	}
}

func GetOperatorPrecedence(token lexer.Token) int {
	switch token.StrVal {
	case "+", "-":
		return 1
	case "/", "*":
		return 2
	case "^":
		return 3
	default:
		return 0
	}
}

func (p *ParserContext) ParseLiteral() SyntaxNode {
	t, _ := p.matchToken(lexer.Literal)
	val, _ := strconv.ParseFloat(t.StrVal, 64)
	return LiteralNode{val}
}

func CreateUnaryOperation(t lexer.Token, s SyntaxNode) SyntaxNode {
	switch t.StrVal {
	case "+":
		return PlusUnaryOperationNode{s}
	case "-":
		return MinusUnaryOperationNode{s}
	default:
		panic("Not implemented")
	}
}

func CreateBinaryOperation(op lexer.Token, left SyntaxNode, right SyntaxNode) SyntaxNode {
	switch op.StrVal {
	case "+":
		return PlusBinaryOperationNode{left, right}
	case "-":
		return MinusBinaryOperationNode{left, right}
	case "*":
		return MultiplyBinaryOperationNode{left, right}
	case "/":
		return DevideBinaryOperationNode{left, right}
	case "^":
		return PowerBinaryOperationNode{left, right}
	default:
		panic(fmt.Sprintf("Not implemented for: %v", op.StrVal))
	}
}

type ParserDiagnostic struct {
	Message string
	Span    lexer.Span
}

func (d ParserDiagnostic) Error() string {
	return fmt.Sprintf("ERROR: %v Loc:(%v:%v)", d.Message, d.Span.Start, d.Span.End)
}
