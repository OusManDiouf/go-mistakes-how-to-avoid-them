//go:build !integration

// you should know that running test with a specific tag includes both the files
// without tags and the files matching this tag.

// if we want to include this test file only if integration tests is not enabled
// then negate the targeted test, here it's integration tests
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContract(t *testing.T) {
	assert.True(t, true)
}

func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test")
	}
}
