package day05_test

import (
	day05 "aoc2024/05"
	"aoc2024/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input.txt")
	assert.NoError(t, err)

	res, err := day05.PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 143, res)
}

func TestPartTwo(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input.txt")
	assert.NoError(t, err)

	res, err := day05.PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 123, res)
}
