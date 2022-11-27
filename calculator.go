package calculator

import "fmt"

type Calculator struct {
	preExpression    string
	suffixExpression string
	letterExpression string
	operator         *Stack
}

func NewCalculator() *Calculator {
	return &Calculator{
		operator: NewStack(),
	}
}

func (c *Calculator) ToExpression(str string) (string, error) {
	c.preExpression = str
	c.suffixExpression = ""
	c.letterExpression = ""
	for i := range str {
		s := string(str[i])
		switch {
		case isDigit(s):
			c.suffixExpression += s
			if i+1 >= len(str) || !isDigit(string(str[i+1])) {
				c.suffixExpression += ","
			}
		case isAlpha(s):
			c.letterExpression += s
			if i+1 < len(str) && str[i+1] == '(' { // 函数
				if err := c.dealFuncExpression(); err != nil {
					return "", err
				}
			}
		case s == ",":
			if err := c.popUntilBracket(); err != nil {
				return "", err
			}
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

func (c *Calculator) Cal() (interface{}, error) {
	if c.suffixExpression == "" {
		return nil, fmt.Errorf("还未设置表达式")
	}
	c.letterExpression = ""
	var t float64 = 0
	number := NewStack()
	for i := range c.suffixExpression {
		ch := c.suffixExpression[i]
		switch {
		case isDigit(string(ch)):
			t = t*10 + float64(ch) - '0'
		case ch == ',':
			number.Push(t)
			t = 0
		case isAlpha(string(ch)):
			c.letterExpression += string(ch)
			if i+1 < len(c.suffixExpression) && c.suffixExpression[i+1] == ':' {
				handler, ok := OpHandler[c.letterExpression]
				if !ok || number.Len() < 2 {
					return nil, fmt.Errorf("错误的表达式")
				}
				y := number.Pop().(float64)
				x := number.Pop().(float64)
				v, err := handler(x, y)
				if err != nil {
					return nil, err
				}
				number.Push(v)
				c.letterExpression = ""
			}
		case isOperator(string(ch)):
			if number.Len() < 2 {
				return nil, fmt.Errorf("错误的表达式")
			}
			y := number.Pop().(float64)
			x := number.Pop().(float64)
			v, err := OpHandler[string(ch)](x, y)
			if err != nil {
				return nil, err
			}
			number.Push(v)
		}
	}
	if number.Len() != 1 {
		return nil, fmt.Errorf("错误的表达式")
	}
	return number.Top(), nil
}

func (c *Calculator) SetExpression(expression string) {
	c.suffixExpression = expression
}
