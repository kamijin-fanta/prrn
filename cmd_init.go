package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

func cmdInit(c *cli.Context) error {
	baseDir := c.String("dir")

	err := os.MkdirAll(historiesBase(baseDir), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(migrationsBase(baseDir), os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(mainSqlFile(baseDir), os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	f.Close()

	return nil
}
