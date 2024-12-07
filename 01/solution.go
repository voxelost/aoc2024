package day_01

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
	lines := strings.Split(input, "\n")

	var firstList, secondList []int
	for _, line := range lines {
		lineFields := strings.Fields(line)
		firstS, secondS := lineFields[0], lineFields[1]
		first, err := strconv.Atoi(firstS)
		if err != nil {
			return 0, err
		}
		second, err := strconv.Atoi(secondS)
		if err != nil {
			return 0, err
		}

		firstList = append(firstList, first)
		secondList = append(secondList, second)
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	var res int
	for i, first := range firstList {
		res += int(math.Abs(float64(first - secondList[i])))
	}

	return res, nil
}

func PartTwo(input string) (int, error) {
	lines := strings.Split(input, "\n")
	var firstList []int
	secondListOccurences := make(map[int]int)

	for _, line := range lines {
		lineFields := strings.Fields(line)
		firstS, secondS := lineFields[0], lineFields[1]
		first, err := strconv.Atoi(firstS)
		if err != nil {
			return 0, err
		}
		second, err := strconv.Atoi(secondS)
		if err != nil {
			return 0, err
		}

		firstList = append(firstList, first)
		if _, ok := secondListOccurences[second]; ok {
			secondListOccurences[second]++
		} else {
			secondListOccurences[second] = 1
		}
	}

	var res int

	for _, first := range firstList {
		res += first * secondListOccurences[first]
	}

	return res, nil
}
