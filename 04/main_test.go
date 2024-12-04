package day04_test

import (
	day04 "aoc2024/04"
	"aoc2024/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input.txt")
	assert.NoError(t, err)

	res, err := day04.PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 18, res)
}

func TestPartTwo(t *testing.T) {
	input, err := internal.ReadFileAsString("data/input.txt")
	assert.NoError(t, err)

	res, err := day04.PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 9, res)
}
