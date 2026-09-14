package happynum

import "runtime"

var unhappyMarkers [244]bool

func init() {
	for _, m := range [8]int{89, 145, 42, 37, 58, 20, 4, 16} {
		unhappyMarkers[m] = true
	}
}

var squareSums [1000]int

func init() {
	squareSums[0] = 0
	for i := 1; i < len(squareSums); i++ {
		squareSums[i] = (i%10)*(i%10) + squareSums[i/10]
	}
}

func squareSum(n int) int {
	ss := 0
	val := n

	for val > 0 {
		digit := val % 10
		ss += digit * digit
		val /= 10
	}
	return ss
}

// IsHappy returns `true` when `n` is a happy number
//
//nolint:gocyclo
func IsHappy(n int) bool {
	ss := squareSum(n)

	for ss > 243 {
		ss = squareSum(ss)
	}

	for ss != 1 && !unhappyMarkers[ss] {
		ss = squareSums[ss]
	}
	return ss == 1
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
// in the range `start` -> `end` using a single-threaded approach
func DistinctHappyRangeCount(start, end int) int {
	var total int
	for i := start; i <= end; i++ {
		if isFirstIteration(i) && IsHappy(i) {
			total++
		}
	}
	return total
}

// DistinctHappyRangeCountParallel returns a count of the distinct happy numbers found
// in the range `1` -> `n` using multiple goroutines and channels to maximize CPU usage
func DistinctHappyRangeCountParallel(n int) int {
	if n <= 0 {
		return 0
	}

	numWorkers := runtime.NumCPU()
	chunkSize := (n + numWorkers - 1) / numWorkers

	results := make(chan int, numWorkers)

	for w := range numWorkers {
		start := w*chunkSize + 1
		end := min(start+chunkSize-1, n)
		if start > n {
			start = n + 1
		}

		go func(start, end int) {
			results <- DistinctHappyRangeCount(start, end)
		}(start, end)
	}

	total := 0
	for range numWorkers {
		total += <-results
	}

	close(results)
	return total
}
