package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var count int
	app := &cli.App{
		Name:                   "greet",
		Usage:                  "fight the loneliness!",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "foo",
				Usage:   "foo greeting",
				Aliases: []string{"f"},
				Count:   &count,
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println("count", count)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
