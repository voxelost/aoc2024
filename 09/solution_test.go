package day_09_test

import (
	day09 "aoc2024/09"
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
			expected:  1928,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  6435922584968,
		},
	}

	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day09.PartOne(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			inputFile: "testdata/test_input.txt",
			expected:  2858,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  0,
		},
	}
	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day09.PartTwo(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}