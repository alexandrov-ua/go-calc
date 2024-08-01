package lexer

import (
	"testing"
)

func Test_Positive(t *testing.T) {
	l := Create("")
	if r := l.Current.Type; r != EndOfFile {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
}

func Test_Literal_Positive(test *testing.T) {
	l := Create("12")
	if t, v := l.Current.Type, l.Current.StrVal; t != Literal || v != "12" {
		test.Errorf("Error, got: %v and %v, want: %v and %v", t, v, Literal, "12")
	}
	l.MoveNext()
	if r := l.Current.Type; r != EndOfFile {
		test.Errorf("Result was incorrect, got: %v, want: %v.", r, EndOfFile)
	}
}

func Test_Literal_Op_Literal(t *testing.T) {
	l := Create("1+2")
	if r := l.Current.Type; r != Literal {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Operator {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Operator)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Literal {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
	l.MoveNext()
	if r := l.Current.Type; r != EndOfFile {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, EndOfFile)
	}
}

func Test_Literal_Op_Literal_Space(t *testing.T) {
	l := Create(" 1 + 2 ")
	if r := l.Current.Type; r != Literal {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Operator {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Operator)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Literal {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
	l.MoveNext()
	if r := l.Current.Type; r != EndOfFile {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, EndOfFile)
	}
}

func Test_Literal_Op_Literal_Space_Return(t *testing.T) {
	l := Create(" 1 +\n2")
	if r := l.Current.Type; r != Literal {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Operator {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Operator)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Literal {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
	l.MoveNext()
	if r := l.Current.Type; r != EndOfFile {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, EndOfFile)
	}
}

func Test_Literal_Op_Literal_Space_Parentheses(t *testing.T) {
	l := Create(" (1 + 2) ^ 3 ")
	if r := l.Current.Type; r != OpenParentheses {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, OpenParentheses)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Literal {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Operator {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Operator)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Literal {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
	l.MoveNext()
	if r := l.Current.Type; r != CloseParentheses {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, CloseParentheses)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Operator {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Operator)
	}
	l.MoveNext()
	if r := l.Current.Type; r != Literal {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, Literal)
	}
	l.MoveNext()
	if r := l.Current.Type; r != EndOfFile {
		t.Errorf("Result was incorrect, got: %v, want: %v.", r, EndOfFile)
	}
}
