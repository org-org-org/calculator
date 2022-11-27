package calculator

import "fmt"

func (c *Calculator) dealOperator(s string, index int) error {
	// 负数开头的特殊情况：是第一个数，或前一个操作符是括号
	if s == "-" && (index == 0 || c.preExpression[index-1] == '(') {
		c.expression += "0,"
	}
	for !c.operator.Empty() && priority[s] <= priority[c.operator.Top().(string)] {
		err := c.popOperator()
		if err != nil {
			return err
		}
	}
	c.operator.Push(s)
	return nil
}

func (c *Calculator) dealRightBracket() error {
	for !c.operator.Empty() && c.operator.Top() != "(" {
		err := c.popOperator()
		if err != nil {
			return err
		}
	}
	if c.operator.Empty() {
		return fmt.Errorf("表达式错误，请检查是否有多余的右括号")
	}
	c.operator.Pop()
	if !c.operator.Empty() && len(c.operator.Top().(string)) > 1 {
		err := c.popOperator()
		if err != nil {
			return err
		}
	}
	return nil
}
