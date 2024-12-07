package day_05_test

import (
	day05 "aoc2024/05"
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
			expected:  143,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  6267,
		},
	}

	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day05.PartOne(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			inputFile: "testdata/test_input.txt",
			expected:  123,
		},
		{
			inputFile: "testdata/input.txt",
			expected:  5184,
		},
	}
	for _, tc := range testCases {
		input, err := internal.ReadFileAsString(tc.inputFile)
		assert.NoError(t, err)

		res, err := day05.PartTwo(input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}
