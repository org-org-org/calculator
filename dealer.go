package calculator

import (
	"fmt"
	"strconv"
)

func (c *Calculator) popOperator() error {
	op := c.operator.Pop().(string)
	if op == "(" {
		return fmt.Errorf("表达式错误，请检查是否有多余的左括号")
	}
	if len(op) > 1 {
		op += "@"
	}
	c.suffixExpression += op
	return nil
}

func (c *Calculator) dealFuncExpression() error {
	_, ok := OpHandler[c.letters]
	if !ok {
		return fmt.Errorf("未知的函数%s", c.letters)
	}
	c.operator.Push(c.letters + ":1")
	c.letters = ""
	return nil
}

func (c *Calculator) dealOperator(s string, index int) error {
	// 负数开头的特殊情况：是第一个数，或前一个操作符是括号
	if s == "-" && (index == 0 || c.preExpression[index-1] == '(') {
		c.suffixExpression += "0,"
	}
	for !c.operator.Empty() && priority[s] <= priority[c.operator.Top().(string)] {
		if err := c.popOperator(); err != nil {
			return err
		}
	}
	c.operator.Push(s)
	return nil
}

func (c *Calculator) addArgsCnt() error {
	str, i, index, err := c.operator.LastContains(':')
	if err != nil {
		return err
	}
	cnt, err := strconv.Atoi(str[index+1:])
	if err != nil {
		return err
	}
	return c.operator.SetIndex(i, str[:index+1]+strconv.Itoa(cnt+1))
}

func (c *Calculator) popUntilBracket() error {
	for !c.operator.Empty() && c.operator.Top() != "(" {
		if err := c.popOperator(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calculator) dealRightBracket() error {
	if err := c.popUntilBracket(); err != nil {
		return err
	}
	if c.operator.Empty() {
		return fmt.Errorf("表达式错误，请检查是否有多余的右括号")
	}
	c.operator.Pop()
	if !c.operator.Empty() && len(c.operator.Top().(string)) > 1 {
		return c.popOperator()
	}
	return nil
}

func (c *Calculator) calFunction(i, next int, number *Stack) error {
	cnt, err := strconv.Atoi(c.suffixExpression[i+2 : i+2+next])
	if err != nil {
		return err
	}
	handler, ok := OpHandler[c.letters]
	if !ok || number.Len() < cnt {
		return fmt.Errorf("错误的表达式")
	}
	c.letters = ""
	args := make([]float64, cnt)
	for j := cnt - 1; j >= 0; j-- {
		args[j] = number.Pop().(float64)
	}
	v, err := handler(args...)
	if err != nil {
		return err
	}
	number.Push(v)
	return nil
}
