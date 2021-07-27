package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
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
				&cli.StringFlag{
					Name:        "tool",
					Required:    false,
					DefaultText: "sql-migrate",
					Usage:       "supported " + strings.Join([]string{SQLMigrate, GolangMigrate}, ", "),
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
