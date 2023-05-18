package happynum

import (
	"math"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func squareSum(n int) int {
	ss := 0
	val := 0 + n

	for val > 0 {
		ss += int(math.Pow(float64(val%10), 2))
		val = int(val / 10)
	}
	return ss
}

// IsHappy returns `true` when `n` is a happy number
func IsHappy(n int) bool {
	unhappyMarkers := []int{89, 145, 42, 37, 58, 20, 4, 16}
	ss := n

	for {
		if ss == 1 {
			return true
		} else if contains(unhappyMarkers, ss) {
			return false
		} else {
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
