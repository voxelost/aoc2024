package main

import (
	day01 "aoc2024/01"
	"aoc2024/internal"
	"fmt"
	"log/slog"
)

func main() {
	input, err := internal.ReadFileAsString("01/data/input.txt")
	if err != nil {
		slog.Error(err.Error())
	}

	res, err := day01.PartOne(input)
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Part one result: %d", res))

	res, err = day01.PartTwo(input)
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Part one result: %d", res))
}
