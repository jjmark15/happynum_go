package main

import "fmt"

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
			ss := ss + 1
		}
	}
}

func main() {
	fmt.Println(isHappy(10))
}
