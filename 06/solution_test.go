package day_06_test

import (
	day06 "aoc2024/06"
	"aoc2024/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	inputFile string
	expected  int
}

func TestPartOne(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			inputFile: "testdata/test_input.txt",
			expected:  41,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  4665,
		},
	}

	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day06.PartOne(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			inputFile: "testdata/test_input.txt",
			expected:  6,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  1688,
		},
	}
	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day06.PartTwo(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}
