package calculator

import (
	"fmt"
	"math"
)

var OpHandler map[string]func(...float64) (float64, error)
var priority = map[string]int8{
	"+":   10,
	"-":   10,
	"*":   20,
	"/":   20,
	"^":   30,
	"(":   0,
	"MAX": 100,
	"MIN": 100,
}

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

	OpHandler["MAX"] = func(v ...float64) (float64, error) {
		if len(v) == 0 {
			return 0, fmt.Errorf("参数错误")
		}
		ans := -2e9
		for i := range v {
			if v[i] > ans {
				ans = v[i]
			}
		}
		return ans, nil
	}

	OpHandler["MIN"] = func(v ...float64) (float64, error) {
		if len(v) == 0 {
			return 0, fmt.Errorf("参数错误")
		}
		ans := 2e9
		for i := range v {
			if v[i] < ans {
				ans = v[i]
			}
		}
		return ans, nil
	}
}
