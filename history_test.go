package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHistory(t *testing.T) {
	assert := assert.New(t)

	h := &history{
		Id:   123,
		Name: "init",
	}
	assert.Equal("000123_init", h.String())

	res, err := parseHistory("000123_init")
	assert.Nil(err)
	assert.Equal(&history{
		Id:   123,
		Name: "init",
	}, res)
}
