package math

func Pow(x, n int64) int64 {
	var acc int64 = 1
	if n < 0 {
		panic("n must not be negative")
	}
	for {
		if n == 0 {
			return acc
		} else if n%2 == 0 {
			x = x * x
			n = n >> 1
		} else {
			acc = acc * x
			n--
		}
	}
}

func Powf(x float64, n int64) float64 {
	var acc float64 = 1
	if n < 0 {
		panic("n must not be negative")
	}
	for {
		if n == 0 {
			return acc
		} else if n%2 == 0 {
			x = x * x
			n = n >> 1
		} else {
			acc = acc * x
			n--
		}
	}
}
