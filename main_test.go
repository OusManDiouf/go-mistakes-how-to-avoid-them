package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCountEmptyLineInFileV1(t *testing.T) {

	count, err := countEmptyLineInFileV1("file.txt")
	require.NoError(t, err)
	assert.Equal(t, 5, count, "count not match")
}

func countEmptyLineInFileV1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	count := 0
	sc := bufio.NewScanner(file) //by default, split the input per line
	for sc.Scan() {

		if sc.Text() == "" {
			//if len(sc.Text()) == 0 {
			//if len(sc.Bytes()) == 0 {
			//if utf8.RuneCountInString(sc.Text()) == 0 {
			count += 1
		}
	}
	return count, err
}
