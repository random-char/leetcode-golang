package solutions

func MySqrt(x int) int {
	return mySqrt(x)
}

func mySqrt(x int) int {
	l, r := 1, max(x/2, 1)
	var mid, sqr int

	for l <= r {
		mid = (l + r) / 2
		sqr = mid * mid
		if sqr == x {
			return mid
		} else if sqr > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return min(l, r)
}
