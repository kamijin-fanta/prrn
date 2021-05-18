package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	godotenv.Load()

	app := &cli.App{
		Name:  "prrn",
		Usage: "simple sql migrator",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Value:   "schema",
				EnvVars: []string{"PRRN_DIR"},
			},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:   "init",
			Action: cmdInit,
		},
		{
			Name: "make",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Required: true,
				},
			},
			Action: cmdMake,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
