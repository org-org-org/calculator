package calculator

import "fmt"

func (c *Calculator) popOperator() error {
	op := c.operator.Pop()
	if op == "(" {
		return fmt.Errorf("表达式错误，请检查是否有多余的左括号")
	}
	c.expression += op.(string)
	return nil
}

func isDigit(str string) bool {
	for i := range str {
		if str[i] < '0' || str[i] > '9' {
			return false
		}
	}
	return true
}

func isAlpha(str string) bool {
	for i := range str {
		if !isLower(str[i]) && !isUpper(str[i]) {
			return false
		}
	}
	return true
}

func isLower(ch uint8) bool {
	return ch >= 'a' && ch <= 'z'
}

func isUpper(ch uint8) bool {
	return ch >= 'A' && ch <= 'Z'
}

func (c *Calculator) SetExpression(expression string) {
	c.expression = expression
}

func (c *Calculator) isOperator(str string) bool {
	_, ok := priority[str]
	return ok
}
