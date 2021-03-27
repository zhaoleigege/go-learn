package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var lang string
	app := &cli.App{
		Name:  "test",
		Usage: "test test",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Usage:       "语言",
				Value:       "english",
				Destination: &lang,
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Printf("action执行: %s, ctx-lang: %s, lang: %s\n",
				ctx.Args().Get(0), ctx.String("lang"), lang)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
