package day_03

import (
	"regexp"
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
	input = strings.ReplaceAll(input, "\n", "")

	re, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		return 0, err
	}

	var acc int
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}

		acc += a * b
	}

	return acc, nil
}

func PartTwo(input string) (int, error) {
	input = strings.ReplaceAll(input, "\n", "")

	re, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	if err != nil {
		return 0, err
	}

	matches, err := re.FindAllStringSubmatch(input, -1), nil
	if err != nil {
		return 0, err
	}

	var acc int
	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
			continue
		} else if match[0] == "don't()" {
			enabled = false
			continue
		}

		if enabled {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, err
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				return 0, err
			}

			acc += a * b
		}
	}

	return acc, nil
}
