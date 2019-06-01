package happynum

import (
	"math"
	"strconv"
	"strings"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func squareSum(n *int) int {
	strN := strconv.Itoa(*n)
	var total int
	for _, v := range strings.Split(strN, "") {
		if intV, err := strconv.Atoi(v); err == nil {
			total += int(math.Pow(float64(intV), 2))
		}
	}
	return total
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
			ss = squareSum(&ss)
		}
	}
}

func isFirstIteration(n int) bool {
	strN := strconv.Itoa(n)
	prev := rune('0')

	for i, currRune := range strN {
		curr := currRune
		if i > 0 {
			if prev > curr {
				return false
			}
		}
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
