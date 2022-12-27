package cli

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jjmark15/happynum_go/pkg/happynum"
	"github.com/urfave/cli"
)

var tagVersion string

func interpretArg(argS string) int {
	if arg, err := strconv.ParseFloat(argS, 64); err == nil {
		return int(arg)
	}
	return 1
}

// Run returns an instance of a urfave cli
func Run() {
	var checkRange string

	app := cli.NewApp()

	app.Name = "Distinct Happy Number Range Counter"
	app.Version = tagVersion
	app.Authors = []cli.Author{
		{
			Name:  "Josh Jones",
			Email: "ohblonddev@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "range, r",
			Value:       "1e6",
			Usage:       "`RANGE` to be calculated",
			Destination: &checkRange,
		},
	}

	app.Action = func(c *cli.Context) error {
		start := time.Now()

		found := happynum.DistinctHappyRangeCount(interpretArg(checkRange))

		elapsed := time.Since(start)
		fmt.Printf("%d distinct happy numbers found in %s\n", found, elapsed)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
