package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jjmark15/happynum_go/happynum"
)

func getArg() int {
	if len(os.Args) > 1 {
		return interpretArg(os.Args[1])
	}
	return 1
}

func interpretArg(argS string) int {
	if arg, err := strconv.ParseFloat(argS, 64); err == nil {
		return int(arg)
	}
	return 1
}

// CliTool happy number cli tool
func cliTool() {
	arg := getArg()
	start := time.Now()

	found := happynum.DistinctHappyRangeCount(arg)

	elapsed := time.Since(start)
	fmt.Printf("%d distinct happy numbers found in %s\n", found, elapsed)
}

func main() {
	cliTool()
}
