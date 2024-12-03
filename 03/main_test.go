package day03_test

import (
	day03 "aoc2024/03"
	"aoc2024/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input_1.txt")
	assert.NoError(t, err)

	res, err := day03.PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 161, res)
}

func TestPartTwo(t *testing.T) {
	input, err := internal.ReadFileAsString("data/test_input_2.txt")
	assert.NoError(t, err)

	res, err := day03.PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 48, res)
}
