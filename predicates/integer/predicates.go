package integer

func NotZero(i int) bool {
	return i != 0
}

func PositiveInt(i int) bool {
	return i > 0
}

func InBetween(min int) func(int) func(int) bool {
	return func(max int) func(int) bool {
		return func(i int) bool {
			return min <= i && i <= max
		}
	}
}
