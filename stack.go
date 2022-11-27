package calculator

import (
	"fmt"
	"strings"
)

type Stack struct {
	a []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	return len(s.a)
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(v ...interface{}) {
	s.a = append(s.a, v...)
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	v := s.a[s.Len()-1]
	s.a = s.a[:s.Len()-1]
	return v
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	return s.a[s.Len()-1]
}

func (s *Stack) LastContains(ch uint8) (string, int, int, error) {
	for i := s.Len() - 1; i >= 0; i-- {
		str, ok := s.a[i].(string)
		if !ok {
			return "", i, 0, fmt.Errorf("不是字符串类型")
		}
		index := strings.IndexByte(str, ch)
		if index != -1 {
			return str, i, index, nil
		}
	}
	return "", 0, 0, fmt.Errorf("没有找到包含%v的元素", ch)
}

func (s *Stack) SetIndex(i int, v interface{}) error {
	if i < 0 || i >= s.Len() {
		return fmt.Errorf("错误的下标")
	}
	s.a[i] = v
	return nil
}
