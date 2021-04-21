package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func getLastHistory(baseDir string) (*history, []byte, error) {
	historiesBase := filepath.Join(baseDir, "histories")
	histories, err := getSortedHistories(historiesBase) // todo configurable

	if err != nil {
		return nil, nil, err
	}
	if len(histories) == 0 {
		return &history{}, []byte{}, nil
	}

	last := histories[len(histories)-1]
	content, err := readFileAll(filepath.Join(historiesBase, last.SqlFilename()))
	if err != nil {
		return nil, nil, err
	}
	return last, content, nil
}

func getSortedHistories(historyPath string) ([]*history, error) {
	files, err := ioutil.ReadDir(historyPath)
	if err != nil {
		return nil, err
	}

	histories := make([]*history, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}

		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		history, err := parseHistory(name)
		if err != nil {
			log.Printf("invalid file name %s", file.Name())
			continue
		}

		histories = append(histories, history)
	}

	sort.Slice(histories, func(i, j int) bool {
		return histories[i].Id < histories[j].Id
	})

	return histories, nil
}

func parseHistory(name string) (*history, error) {
	delm := strings.Index(name, "_")
	if delm == -1 {
		return nil, fmt.Errorf("invalid input format")
	}
	id, err := strconv.Atoi(name[0:delm])
	if err != nil {
		return nil, err
	}

	return &history{
		Id:   id,
		Name: name[delm+1:],
	}, nil
}

type history struct {
	Id   int
	Name string
}

func (h *history) String() string {
	return fmt.Sprintf("%06d_%s", h.Id, h.Name)
}

func (h *history) SqlFilename() string {
	return h.String() + ".sql"
}

