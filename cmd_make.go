package main

import (
	"bytes"
	"fmt"
	"github.com/schemalex/schemalex/diff"
	"github.com/urfave/cli/v2"
)

const (
	SQLMigrate    = "sql-migrate"
	GolangMigrate = "golang-migrate"
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

	in := migrateInput{
		name:    newName,
		baseDir: baseDir,
		history: last,
		lastSQL: lastSql,
		mainSQL: mainSql,
	}

	migrationTool := c.String("tool")
	switch migrationTool {
	case SQLMigrate:
		return makeMigrationFileSqlMigrate(in)
	case GolangMigrate:
		return makeMigrationFileGoMigrate(in)
	default: // default use sql-migrate
		return makeMigrationFileSqlMigrate(in)
	}
}

type migrateInput struct {
	name    string
	baseDir string

	history *history
	lastSQL []byte
	mainSQL []byte
}

func makeMigrationFileSqlMigrate(in migrateInput) error {
	h := &history{
		Id:   in.history.Id + 1,
		Name: in.name,
	}
	migrationSql := &bytes.Buffer{}
	migrationSql.Write([]byte("-- +migrate Up\nSET FOREIGN_KEY_CHECKS = 0;\n"))
	err := diff.Strings(migrationSql, string(in.lastSQL), string(in.mainSQL))
	if err != nil {
		return err
	}

	migrationSql.Write([]byte("\nSET FOREIGN_KEY_CHECKS = 1;\n\n\n-- +migrate Down\nSET FOREIGN_KEY_CHECKS = 0;\n"))
	err = diff.Strings(migrationSql, string(in.mainSQL), string(in.lastSQL))
	if err != nil {
		return err
	}

	migrationSql.Write([]byte("\nSET FOREIGN_KEY_CHECKS = 1;\n"))
	if err := putMigrationFile(in.baseDir, fmt.Sprintf("%s.sql", h.String()), migrationSql.Bytes()); err != nil {
		return err
	}
	return putHistory(in.baseDir, h, in.mainSQL)
}

func makeMigrationFileGoMigrate(in migrateInput) error {
	h := &history{
		Id:   in.history.Id + 1,
		Name: in.name,
	}
	migrationSql := &bytes.Buffer{}
	err := diff.Strings(migrationSql, string(in.lastSQL), string(in.mainSQL))
	if err != nil {
		return err
	}
	err = putMigrationFile(in.baseDir, fmt.Sprintf("%s.up.sql", h.String()), migrationSql.Bytes())
	if err != nil {
		return err
	}

	err = diff.Strings(migrationSql, string(in.mainSQL), string(in.lastSQL))
	if err != nil {
		return err
	}

	err = putMigrationFile(in.baseDir, fmt.Sprintf("%s.down.sql", h.String()), migrationSql.Bytes())
	if err != nil {
		return err
	}

	return putHistory(in.baseDir, h, in.mainSQL)
}
