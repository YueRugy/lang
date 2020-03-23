package calc

func Add(x, y int) int {
	return x + y
}

func Del(x, y int) int {
	return x - y
}

func Da(x, y int) int {
	return x * y
}

func Jc(x int) int {
	if x < 1 {
		return 0
	}

	if x == 1 {
		return 1
	}

	return Jc(x-1) * x
}
