package day06_test

import (
	day06 "aoc2024/06"
	"aoc2024/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input.txt")
	assert.NoError(t, err)

	res, err := day06.PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 41, res)
}

func TestPartTwo(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input.txt")
	assert.NoError(t, err)

	res, err := day06.PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 6, res)
}
