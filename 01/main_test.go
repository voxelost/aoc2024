package day01_test

import (
	day01 "aoc2024/01"
	"aoc2024/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input.txt")
	assert.NoError(t, err)

	res, err := day01.PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 11, res)
}

func TestPartTwo(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input.txt")
	assert.NoError(t, err)

	res, err := day01.PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 31, res)
}
