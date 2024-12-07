package day_02_test

import (
	day02 "aoc2024/02"
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
			expected:  2,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  572,
		},
	}

	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day02.PartOne(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			inputFile: "testdata/test_input.txt",
			expected:  4,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  612,
		},
	}
	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day02.PartTwo(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}
