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

func putHistory(baseDir string, history *history, full []byte) error {
	return os.WriteFile(filepath.Join(historiesBase(baseDir), history.SqlFilename()), full, os.ModePerm)

}

func putMigrationFile(baseDir string, sqlFileName string, migration []byte) error {
	return os.WriteFile(filepath.Join(migrationsBase(baseDir), sqlFileName), migration, os.ModePerm)
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
