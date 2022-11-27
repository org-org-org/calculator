package tests

import (
	"calculator"
	"fmt"
	"testing"
)

func TestCal(t *testing.T) {
	cal := calculator.NewCalculator()
	s, err := cal.ToExpression("(2/4)^(-3+3)")
	fmt.Println(s)
	if err != nil {
		t.Error(err)
	}
	ex, err := cal.Cal()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ex)
}
