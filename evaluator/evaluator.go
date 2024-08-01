package evaluator

import (
	"calc/parser"
	"errors"
)

func Evaluate(s string) (float64, error) {
	st, diag := parser.Parse(s)

	if len(diag) > 0 {
		errs := make([]error, len(diag))
		for i, d := range diag {
			errs[i] = d
		}

		return 0, errors.Join(errs...)
	}

	return st.Exec(), nil
}
