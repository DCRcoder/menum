package main

import (
	"os"

	"github.com/DCRcoder/menum/generator"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:            "enumer",
		Usage:           "An enum generator for go",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "The file(s) to generate enums.  Use more than one flag for more files.",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Usage:    "output file name; default srcdir/<type>_enum.go",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			generator.GeneratorFunc(ctx.String("file"), ctx.String("output"))
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
