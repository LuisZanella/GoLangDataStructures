package main

func mySqrt(x int) (result int) {
	if x <= 0 {
		return 0
	}
	l := 1
	r := x
	for l <= r {
		m := l + (r-l)/2
		if m*m == x {
			return m
		}
		if m*m > x {
			r = m - 1
		} else {
			result = m
			l = m+1
		}

	}
	return result
}
