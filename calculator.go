package calculator

import "fmt"

type Calculator struct {
	preExpression string
	operator      *Stack
	expression    string
}

func NewCalculator() *Calculator {
	return &Calculator{
		operator: NewStack(),
	}
}

func (c *Calculator) ToExpression(str string) (string, error) {
	c.preExpression = str
	c.expression = ""
	alpha := ""
	for i := range c.preExpression {
		s := string(str[i])
		switch {
		case isDigit(s):
			c.expression += s
			if i+1 >= len(str) || !isDigit(string(str[i+1])) {
				c.expression += ","
			}
		case isAlpha(s):
			alpha += s
			if i+1 < len(str) && str[i+1] == '(' {
				//if err := c.dealOperator(alpha+":", i); err != nil {
				//	return "", err
				//}
				_, ok := OpHandler[alpha]
				if ok {
					c.operator.Push(alpha + ":")
				}
				alpha = ""
			}
		case s == ",":
			for !c.operator.Empty() && c.operator.Top() != "(" {
				err := c.popOperator()
				if err != nil {
					return "", err
				}
			}
		case s == "(":
			c.operator.Push(s)
		case s == ")": // 右括号
			if err := c.dealRightBracket(); err != nil {
				return "", err
			}
		case c.isOperator(s): // 操作符
			if err := c.dealOperator(s, i); err != nil {
				return "", err
			}
		}
	}
	for !c.operator.Empty() {
		err := c.popOperator()
		if err != nil {
			return "", err
		}
	}
	return c.expression, nil
}

func (c *Calculator) Cal() (interface{}, error) {
	if c.expression == "" {
		return nil, fmt.Errorf("还未设置表达式")
	}
	var t float64 = 0
	alpha := ""
	number := NewStack()
	for i := range c.expression {
		ch := c.expression[i]
		switch {
		case isDigit(string(ch)):
			t = t*10 + float64(ch) - '0'
		case ch == ',':
			number.Push(t)
			t = 0
		case isAlpha(string(ch)):

			alpha += string(ch)
			if i+1 >= len(c.expression) || c.expression[i+1] == ':' {
				handler, ok := OpHandler[alpha]
				if ok {
					if number.Len() < 2 {
						return nil, fmt.Errorf("错误的表达式")
					}
					y := number.Pop().(float64)
					x := number.Pop().(float64)
					v, err := handler(x, y)
					if err != nil {
						return nil, err
					}
					number.Push(v)
				}
				alpha = ""
			}
		case c.isOperator(string(ch)):
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
