package evaluator

import (
	"testing"
)

func Test_Exec(t *testing.T) {
	res, _ := Evaluate("1+2+3")
	if exp := float64(6); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Exec1(t *testing.T) {
	res, _ := Evaluate("2*3+4")
	if exp := float64(10); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Exec2(t *testing.T) {
	res, _ := Evaluate("2^5*3+4")
	if exp := float64(100); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Exec3(t *testing.T) {
	res, _ := Evaluate("2^(5*3)+4")
	if exp := float64(32772); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Exec4(t *testing.T) {
	res, _ := Evaluate("2^5*(3+4)")
	if exp := float64(224); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Exec5(t *testing.T) {
	res, _ := Evaluate("(((2^5)*3)+4)")
	if exp := float64(100); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Exec6(t *testing.T) {
	res, _ := Evaluate("-(((2^5)*3)+4)")
	if exp := float64(-100); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Exec7(t *testing.T) {
	res, _ := Evaluate("2+-2")
	if exp := float64(0); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Exec8(t *testing.T) {
	res, _ := Evaluate("2++2")
	if exp := float64(4); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Exec9(t *testing.T) {
	res, _ := Evaluate("2^+2")
	if exp := float64(4); res != exp {
		t.Errorf("Expected %v but found %v", exp, res)
	}
}

func Test_Negative(t *testing.T) {
	_, err := Evaluate("***")
	if err == nil {
		t.Errorf("Error is nil")
	} else if err.Error() != "ERROR: Expected Literal but found Operator Loc:(0:1)\nERROR: Expected Literal but found Operator Loc:(2:3)" {
		t.Errorf("Wrong error message: %v", err.Error())
	}
}

func Test_Negative1(t *testing.T) {
	_, err := Evaluate("---")
	if err == nil {
		t.Errorf("Error is nil")
	} else if err.Error() != "ERROR: Expected Literal but found EndOfFile Loc:(2:3)" {
		t.Errorf("Wrong error message: %v", err.Error())
	}
}

func Test_Negative2(t *testing.T) {
	_, err := Evaluate("2**3")
	if err == nil {
		t.Errorf("Error is nil")
	} else if err.Error() != "ERROR: Expected Literal but found Operator Loc:(2:3)\nERROR: Expected Operator but found Literal Loc:(3:4)" {
		t.Errorf("Wrong error message: %v", err.Error())
	}
}

func Test_Negative3(t *testing.T) {
	_, err := Evaluate("sdfsd")
	if err == nil {
		t.Errorf("Error is nil")
	} else if err.Error() != "ERROR: Expected Literal but found Identifier Loc:(0:5)" {
		t.Errorf("Wrong error message: %v", err.Error())
	}
}
