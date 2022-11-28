package calculator

import (
	"fmt"
	"strings"
)

type Calculator struct {
	preExpression    string
	suffixExpression string
	letters          string
	operator         *Stack
	args             map[string]float64
}

func NewCalculator() *Calculator {
	return &Calculator{
		operator: NewStack(),
	}
}

func (c *Calculator) ToSuffixExpression(str string) (string, error) {
	c.preExpression = str
	c.suffixExpression = ""
	c.letters = ""
	for i := range str {
		s := string(str[i])
		switch {
		case isDigit(s):
			c.suffixExpression += s
			if i+1 >= len(str) || !(isDigit(string(str[i+1])) || str[i+1] == '.') {
				c.suffixExpression += ","
			}
		case isAlpha(s):
			c.letters += s
			if i+1 < len(str) && str[i+1] == '(' { // 函数
				if err := c.dealFuncExpression(); err != nil {
					return "", err
				}
			} else if i+1 >= len(str) || !canBeArgs(str[i+1]) { // 变量
				c.suffixExpression += c.letters + ","
				c.letters = ""
			}
		case s == ",":
			if err := c.addArgsCnt(); err != nil {
				return "", err
			}
			if err := c.popUntilBracket(); err != nil {
				return "", err
			}
		case s == ".":
			c.suffixExpression += s
		case s == "(":
			c.operator.Push(s)
		case s == ")": // 右括号
			if err := c.dealRightBracket(); err != nil {
				return "", err
			}
		case isOperator(s): // 操作符
			if err := c.dealOperator(s, i); err != nil {
				return "", err
			}
		}
	}
	for !c.operator.Empty() {
		if err := c.popOperator(); err != nil {
			return "", err
		}
	}
	return c.suffixExpression, nil
}

func (c *Calculator) Cal() (float64, error) {
	if c.suffixExpression == "" {
		return 0, fmt.Errorf("还未设置表达式")
	}
	c.letters = ""
	var t float64 = 0
	number := NewStack()
	isFloat := false
	for i := 0; i < len(c.suffixExpression); i++ {
		ch := c.suffixExpression[i]
		s := string(ch)
		switch {
		case isDigit(s):
			if isFloat {
				t = t + 0.1*(float64(ch)-'0')
			} else {
				t = t*10 + float64(ch) - '0'
			}
		case ch == ',':
			if c.letters != "" {
				if c.args == nil {
					return 0, fmt.Errorf("未设置参数map")
				}
				v, ok := c.args[c.letters]
				if !ok {
					return 0, fmt.Errorf("未知的参数%s", c.letters)
				}
				number.Push(v)
				c.letters = ""
			} else {
				number.Push(t)
				t = 0
				isFloat = false
			}
		case ch == '.':
			isFloat = true
		case isAlpha(s):
			c.letters += s
			if i+1 < len(c.suffixExpression) && c.suffixExpression[i+1] == ':' {
				next := strings.IndexRune(c.suffixExpression[i+2:], '@')
				if err := c.calFunction(i, next, number); err != nil {
					return 0, err
				}
				i = i + 2 + next
			}
		case isOperator(s):
			if number.Len() < 2 {
				return 0, fmt.Errorf("错误的表达式")
			}
			y := number.Pop().(float64)
			x := number.Pop().(float64)
			v, err := OpHandler[s](x, y)
			if err != nil {
				return 0, err
			}
			number.Push(v)
		}
	}
	if number.Len() != 1 {
		return 0, fmt.Errorf("错误的表达式")
	}
	return number.Top().(float64), nil
}

func (c *Calculator) SetSuffixExpression(expression string) *Calculator {
	c.suffixExpression = expression
	return c
}

func (c *Calculator) SetArgs(args map[string]float64) *Calculator {
	c.args = args
	return c
}
