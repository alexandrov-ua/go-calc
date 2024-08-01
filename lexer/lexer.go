package lexer

import (
	"unicode"
)

func Create(s string) *Lexer {
	l := Lexer{Runes: []rune(s)}
	l.MoveNext()
	return &l
}

type Lexer struct {
	Runes   []rune
	Current Token
	index   int64
}

func (l *Lexer) GetAndMooveNext() Token {
	tmp := l.Current
	l.MoveNext()
	return tmp
}

func (l *Lexer) MoveNext() {
	for ; int(l.index) < len(l.Runes); l.index++ {
		currRune := l.Runes[l.index]
		switch {
		case unicode.IsDigit(currRune):
			predicate := func(r rune) bool {
				return unicode.IsDigit(r) || r == '.'
			}
			res := readRunes(l.Runes[l.index:], predicate)
			l.Current = Token{Literal, Span{l.index, l.index + int64(len(res))}, string(res)}
			l.index += int64(len(res))
			return
		case unicode.IsLetter(currRune):
			predicate := func(r rune) bool {
				return unicode.IsLetter(r) || unicode.IsNumber(r)
			}
			res := readRunes(l.Runes[l.index:], predicate)
			l.Current = Token{Identifier, Span{l.index, l.index + int64(len(res))}, string(res)}
			l.index += int64(len(res))
			return
		case unicode.IsSpace(currRune):
			continue
		case currRune == '(':
			l.Current = Token{OpenParentheses, Span{l.index, l.index + 1}, string(currRune)}
			l.index++
			return
		case currRune == ')':
			l.Current = Token{CloseParentheses, Span{l.index, l.index + 1}, string(currRune)}
			l.index++
			return
		case currRune == '+' || currRune == '-' || currRune == '/' || currRune == '*' || currRune == '^':
			l.Current = Token{Operator, Span{l.index, l.index + 1}, string(currRune)}
			l.index++
			return
		default:
			l.Current = Token{Unknown, Span{l.index, l.index + 1}, string(currRune)}
			l.index++
			return
		}
	}

	if l.Current.Type != EndOfFile {
		l.Current = Token{EndOfFile, Span{l.index - 1, l.index}, ""}
	}
}

func readRunes(r []rune, f func(rune) bool) []rune {
	c := 0
	for i := 0; i < len(r) && f(r[i]); i++ {
		c++
	}
	return r[:c]
}

type Token struct {
	Type   TokenType
	Span   Span
	StrVal string
}

type Span struct {
	Start int64
	End   int64
}

type TokenType int

func (t TokenType) String() string {
	switch t {
	case NoToken:
		return "NoToken"
	case Literal:
		return "Literal"
	case Operator:
		return "Operator"
	case Identifier:
		return "Identifier"
	case OpenParentheses:
		return "OpenParentheses"
	case CloseParentheses:
		return "CloseParentheses"
	case EndOfFile:
		return "EndOfFile"
	case Unknown:
		return "Unknown"
	default:
		panic("No implemented")
	}
}

const (
	NoToken  TokenType = -1
	Literal  TokenType = 0
	Operator TokenType = iota
	Identifier
	OpenParentheses
	CloseParentheses
	EndOfFile
	Unknown
)
