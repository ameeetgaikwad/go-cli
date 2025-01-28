package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var language string
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "english",
				Usage:       "language for the greeting `spanish`",
				Aliases:     []string{"l"},
				Destination: &language,
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "amit"

			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}
			// if cCtx.String("lang") == "spanish" {
			// 	fmt.Println("Hola", name)

			// }
			if language == "spanish" {
				fmt.Println("Hola", name)

			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
