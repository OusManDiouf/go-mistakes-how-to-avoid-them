package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestCountEmptyLineInFile(t *testing.T) {

	file, err := os.Open("file.txt")
	require.NoError(t, err)
	count, err := countEmptyLine(file)
	require.NoError(t, err)

	assert.Equal(t, 5, count)
}
func TestCountEmptyLineInHTTPBody(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/testOne", strings.NewReader(
		`Try cooking tart marinateed with ketchup, enameled with vodka.How dead.

			You mark like a parrot.`))
	require.NoError(t, err)

	count, err := countEmptyLine(req.Body)
	require.NoError(t, err)
	assert.Equal(t, 1, count)

}

func TestCountEmptyLineInStringReader(t *testing.T) {

	textWithEmptyLine := strings.NewReader(
		`One must follow the lama in order to synthesise the karma of great thought.

		Corsairs die from madnesses like rough clouds.

		Be ancient for whoever converts, because each has been experienced with courage.
		Countless honors will be lost in minerals like powerdrains in energies

	`)

	count, err := countEmptyLine(textWithEmptyLine)
	require.NoError(t, err)
	assert.Equal(t, 3, count, "count not match")
}
