package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
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
	strN := strconv.Itoa(n)
	total := 0
	for _, v := range strings.Split(strN, "") {
		if intV, err := strconv.Atoi(v); err == nil {
			total += int(math.Pow(float64(intV), 2))
		}
	}
	return total
}

func isHappy(n int) bool {
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
	strN := strconv.Itoa(n)
	digitSlice := []int{}
	for _, v := range strings.Split(strN, "") {
		if intV, err := strconv.Atoi(v); err == nil {
			digitSlice = append(digitSlice, intV)
		}
	}

	return sort.IntsAreSorted(digitSlice)
}

func distinctHappyRangeCount(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		if isFirstIteration(i) && isHappy(i) {
			total++
		}
	}
	return total
}

func main() {
	start := time.Now()

	found := distinctHappyRangeCount(1e6)

	elapsed := time.Since(start)
	fmt.Printf("%d unique happy numbers found in %s\n", found, elapsed)
}
