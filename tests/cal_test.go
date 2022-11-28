package tests

import (
	"calculator"
	"math"
	"testing"
)

var cal = calculator.NewCalculator()

func TestCal(t *testing.T) {
	test(t, "4^(-2+2)", 1)
	test(t, "-12/6+5+2*3", 9)
	test(t, "2.5+0.5", 3)
	test(t, "-0.5-1", -1.5)
	test(t, "MAX(MIN(1,2),MIN(3,4))", 3)
	test(t, "MAX(MIN(0.1,2),MIN(0.3,4))", 0.3)
	test(t, "MAX(MIN(-3,2),MIN((-5/4),4))", -1.25)
	test(t, "MAX(MIN(1,2,0.1),MIN(3,4,0.3))", 0.3)
}

func TestCalWithArgs(t *testing.T) {
	testWithArgs(t, "MIN(apple,why,zc)", map[string]float64{
		"apple": 1,
		"why":   0.5,
		"zc":    2,
	}, 0.5)
}

func test(t *testing.T, str string, ans float64) {
	expression, err := cal.ToSuffixExpression(str)
	//fmt.Println(expression)
	if err != nil {
		t.Error(err)
	}
	res, err := cal.Cal()
	if err != nil {
		t.Error(err)
	}
	if math.Abs(res-ans) > 1e-6 {
		t.Logf("原表达式：%s\n", str)
		t.Logf("转化的表达式：%s\n", expression)
		t.Errorf("错误的答案%f，期望为%f\n", res, ans)
	}
}

func testWithArgs(t *testing.T, str string, args map[string]float64, ans float64) {
	expression, err := cal.ToSuffixExpression(str)
	//fmt.Println(expression)
	if err != nil {
		t.Error(err)
	}
	res, err := cal.SetArgs(args).Cal()
	if err != nil {
		t.Error(err)
	}
	if math.Abs(res-ans) > 1e-6 {
		t.Logf("原表达式：%s\n", str)
		t.Logf("转化的表达式：%s\n", expression)
		t.Logf("参数：%v\n", args)
		t.Errorf("错误的答案%f，期望为%f\n", res, ans)
	}
}
