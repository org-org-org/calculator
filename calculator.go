package calculator

import "fmt"

type Calculator struct {
	priority   map[string]int8
	operator   *Stack
	expression string
}

func NewCalculator() *Calculator {
	priority := map[string]int8{
		"+": 10,
		"-": 10,
		"*": 20,
		"/": 20,
		"^": 30,
		//"(": 0,
	}
	return &Calculator{
		priority:   priority,
		operator:   NewStack(),
		expression: "",
	}
}

func (c *Calculator) ToExpression(str string) (string, error) {
	for i := range str {
		s := string(str[i])
		switch {
		case c.isDigit(s):
			c.expression += s
			if i+1 >= len(str) || !c.isDigit(string(str[i+1])) {
				c.expression += ","
			}
		case s == "(":
			c.operator.Push(s)
		case s == ")":
			for !c.operator.Empty() && c.operator.Top() != "(" {
				err := c.popOperator()
				if err != nil {
					return "", err
				}
			}
			if c.operator.Empty() {
				return "", fmt.Errorf("表达式错误，请检查是否有多余的右括号")
			}
			c.operator.Pop()
		case c.isOperator(s):
			// 负数开头的特殊情况：是第一个数，或前一个操作符是括号
			if s == "-" && (i == 0 || str[i-1] == '(') {
				c.expression += "0,"
			}
			for !c.operator.Empty() && c.priority[s] <= c.priority[c.operator.Top().(string)] {
				err := c.popOperator()
				if err != nil {
					return "", err
				}
			}
			c.operator.Push(s)
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

func (c *Calculator) popOperator() error {
	op := c.operator.Pop()
	if op == "(" {
		return fmt.Errorf("表达式错误，请检查是否有多余的左括号")
	}
	c.expression += op.(string)
	return nil
}

func (c *Calculator) isDigit(str string) bool {
	for i := range str {
		if str[i] < '0' || str[i] > '9' {
			return false
		}
	}
	return true
}

func (c *Calculator) SetExpression(expression string) {
	c.expression = expression
}

func (c *Calculator) isOperator(str string) bool {
	_, ok := c.priority[str]
	return ok
}

func (c *Calculator) Cal() (interface{}, error) {
	if c.expression == "" {
		return nil, fmt.Errorf("还未设置表达式")
	}
	var t float64 = 0
	number := NewStack()
	for i := range c.expression {
		ch := c.expression[i]
		switch {
		case c.isDigit(string(ch)):
			t = t*10 + float64(ch) - '0'
		case ch == ',':
			number.Push(t)
			t = 0
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
