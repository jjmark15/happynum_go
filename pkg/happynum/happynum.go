package happynum

var unhappyMarkers = map[int]bool{
	89: true, 145: true, 42: true, 37: true,
	58: true, 20: true, 4: true, 16: true,
}

func squareSum(n int) int {
	ss := 0
	val := n

	for val > 0 {
		digit := val % 10
		ss += digit * digit
		val = val / 10
	}
	return ss
}

// IsHappy returns `true` when `n` is a happy number
func IsHappy(n int) bool {
	ss := n

	for {
		switch {
		case ss == 1:
			return true
		case unhappyMarkers[ss]:
			return false
		default:
			ss = squareSum(ss)
		}
	}
}

func isFirstIteration(n int) bool {
	rem := n
	prev := 9

	for rem > 0 {
		curr := rem % 10

		if curr > prev {
			return false
		}

		rem /= 10
		prev = curr
	}

	return true
}

// DistinctHappyRangeCount returns a count of the distinct happy numbers found
// in the range `0` -> `n`
func DistinctHappyRangeCount(n int) int {
	var total int
	for i := 1; i <= n; i++ {
		if isFirstIteration(i) && IsHappy(i) {
			total++
		}
	}
	return total
}
