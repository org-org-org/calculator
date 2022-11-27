package calculator

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

func isOperator(str string) bool {
	if len(str) != 1 {
		return false
	}
	_, ok := priority[str]
	return ok
}
