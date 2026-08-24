package slicex

func Sum(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}

func Max(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	max := xs[0]
	for _, x := range xs[1:] {
		if x > max {
			max = x
		}
	}
	return max, true
}
