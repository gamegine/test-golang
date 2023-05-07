package math

// Min of a/b
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Subtract subtracts two integers return -1 if negative
func Subtract(a, b int) (res int) {
	res = a - b
	if res < 0 {
		return -1
	}
	return res
}
