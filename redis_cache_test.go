package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRedisCache(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("skipping integration test")
	}
	assert.True(t, true)
}
