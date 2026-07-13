package cli

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jjmark15/happynum_go/pkg/happynum"
	"github.com/urfave/cli/v3"
)

var tagVersion string

func interpretArg(argS string) int {
	if arg, err := strconv.ParseFloat(argS, 64); err == nil {
		return int(arg)
	}
	return 1
}

func newCommand() *cli.Command {
	var checkRange string
	runSingleThreaded := false

	return &cli.Command{
		Name:    "happynum",
		Usage:   "Distinct Happy Number Range Counter",
		Version: tagVersion,
		Authors: []any{
			"Josh Jones <ohblonddev@gmail.com>",
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "range",
				Aliases:     []string{"r"},
				Value:       "1e6",
				Usage:       "`RANGE` to be calculated",
				Destination: &checkRange,
			},
			&cli.BoolFlag{
				Name:        "single",
				Aliases:     []string{"s"},
				Usage:       "run single threaded",
				Destination: &runSingleThreaded,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			start := time.Now()

			var found int
			if runSingleThreaded {
				found = happynum.DistinctHappyRangeCount(1, interpretArg(checkRange))
			} else {
				found = happynum.DistinctHappyRangeCountParallel(interpretArg(checkRange))
			}

			elapsed := time.Since(start)
			fmt.Printf("count: %d\ntime: %s\n", found, elapsed)
			return nil
		},
	}
}

// Run returns an instance of a urfave cli
func Run() {
	if err := newCommand().Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
