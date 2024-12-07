package day_04_test

import (
	day04 "aoc2024/04"
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
			expected:  18,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  2685,
		},
	}

	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day04.PartOne(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			inputFile: "testdata/test_input.txt",
			expected:  9,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  2048,
		},
	}
	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day04.PartTwo(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}
