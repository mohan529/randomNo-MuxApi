package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Output(t *testing.T) {

	result := randomeNumbers()

	assert.GreaterOrEqual(t, result, 0)
	assert.LessOrEqual(t, result, 1000)
}
