package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSome(t *testing.T) {
	/**
	Because of the overhead, it’s generally recommended to enable the race detector
	only during local testing or continuous integration (CI).
	In production, we should avoid it (or only use it in the case of canary releases, for example).
	*/

	assert.Equal(t, 1, getCount(), "i not matching")
}
