package calculator

import (
	"fmt"
	"math"
)

var OpHandler map[string]func(...float64) (float64, error)

func init() {
	OpHandler = make(map[string]func(...float64) (float64, error), 0)
	OpHandler["+"] = func(v ...float64) (float64, error) {
		if len(v) != 2 {
			return 0, fmt.Errorf("参数错误")
		}
		return v[0] + v[1], nil
	}

	OpHandler["-"] = func(v ...float64) (float64, error) {
		if len(v) != 2 {
			return 0, fmt.Errorf("参数错误")
		}
		return v[0] - v[1], nil
	}

	OpHandler["*"] = func(v ...float64) (float64, error) {
		if len(v) != 2 {
			return 0, fmt.Errorf("参数错误")
		}
		return v[0] * v[1], nil
	}

	OpHandler["/"] = func(v ...float64) (float64, error) {
		if len(v) != 2 {
			return 0, fmt.Errorf("参数错误")
		}
		if v[1] == 0 {
			return 0, fmt.Errorf("除数不能为0")
		}
		return v[0] / v[1], nil
	}

	OpHandler["^"] = func(v ...float64) (float64, error) {
		if len(v) != 2 {
			return 0, fmt.Errorf("参数错误")
		}
		return math.Pow(v[0], v[1]), nil
	}
}
