package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func readFileAll(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return buff, err
}

func put(baseDir string, history *history, migration, full []byte) error {

	err := os.WriteFile(filepath.Join(historiesBase(baseDir), history.SqlFilename()), full, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(migrationsBase(baseDir), history.SqlFilename()), migration, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func historiesBase(baseDir string) string {
	return filepath.Join(baseDir, "histories")
}

func migrationsBase(baseDir string) string {
	return filepath.Join(baseDir, "migrations")
}

func mainSqlFile(baseDir string) string {
	return filepath.Join(baseDir, "main.sql")
}
