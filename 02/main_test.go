package day02_test

import (
	day02 "aoc2024/02"
	"aoc2024/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input.txt")
	assert.NoError(t, err)

	res, err := day02.PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 2, res)
}

func TestPartTwo(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input.txt")
	assert.NoError(t, err)

	res, err := day02.PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 4, res)
}
