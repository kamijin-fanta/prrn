package main

import (
	"bytes"
	"github.com/schemalex/schemalex/diff"
	"github.com/urfave/cli/v2"
)

func cmdMake(c *cli.Context) error {
	baseDir := c.String("dir")
	newName := c.String("name")

	last, lastSql, err := getLastHistory(baseDir)
	if err != nil {
		return err
	}
	mainSql, err := readFileAll(mainSqlFile(baseDir))
	if err != nil {
		return err
	}

	migrationSql := &bytes.Buffer{}
	migrationSql.Write([]byte("-- +migrate Up\nSET FOREIGN_KEY_CHECKS = 0;\n"))
	err = diff.Strings(migrationSql, string(lastSql), string(mainSql))
	if err != nil {
		return err
	}

	migrationSql.Write([]byte("\nSET FOREIGN_KEY_CHECKS = 1;\n\n\n-- +migrate Down\nSET FOREIGN_KEY_CHECKS = 0;\n"))
	err = diff.Strings(migrationSql, string(mainSql), string(lastSql))
	if err != nil {
		return err
	}

	migrationSql.Write([]byte("\nSET FOREIGN_KEY_CHECKS = 1;\n"))
	newHistory := &history{
		Id:   last.Id + 1,
		Name: newName,
	}
	err = put(baseDir, newHistory, migrationSql.Bytes(), mainSql)
	if err != nil {
		return err
	}

	return nil
}
