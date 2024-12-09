package day_07

import (
	"fmt"
	"strconv"
	"strings"
)

type calibration struct {
	lvalue  int64
	rvalues []int64
}

type operation string

var (
	OpAdd    operation = "+"
	OpMul    operation = "*"
	OpConcat operation = "||"
)

func parseInput(input string) ([]calibration, error) {
	lines := strings.Split(input, "\n")
	calibrations := make([]calibration, 0, len(lines))

	for _, line := range lines {
		lineSplit := strings.Fields(line)
		testValueS := strings.TrimSuffix(lineSplit[0], ":")
		equationsS := lineSplit[1:]

		testValue, err := strconv.Atoi(testValueS)
		if err != nil {
			return nil, err
		}

		var equations []int64
		for _, equationS := range equationsS {
			equation, err := strconv.Atoi(equationS)
			if err != nil {
				return nil, err
			}
			equations = append(equations, int64(equation))
		}

		calibrations = append(calibrations, calibration{
			lvalue:  int64(testValue),
			rvalues: equations,
		})
	}

	return calibrations, nil
}

func concatenate(a, b int64) (int64, error) {
	res, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return int64(res), err
}

func getPermutationsOfNOps(opSet []operation, n int) [][]operation {
	if n == 0 {
		return [][]operation{{}}
	}

	var result [][]operation
	for _, prefix := range getPermutationsOfNOps(opSet, n-1) {
		for _, op := range opSet {
			// this made me want to
			p := make([]operation, len(prefix))
			copy(p, prefix)
			result = append(result, append(p, op))
		}
	}
	return result
}

func solveEquation(values []int64, operations []operation) int64 {
	result := values[0]
	for i := 1; i < len(values); i++ {
		switch operations[i-1] {
		case OpAdd:
			result += values[i]
		case OpMul:
			result *= values[i]
		case OpConcat:
			concatenated, err := concatenate(result, values[i])
			if err != nil {
				panic(err)
			}
			result = concatenated
		}
	}
	return result
}

func (c calibration) trySolve(opSet []operation) bool {
	possibleOperations := getPermutationsOfNOps(opSet, len(c.rvalues)-1)

	for _, op := range possibleOperations {
		if c.lvalue == solveEquation(c.rvalues, op) {
			return true
		}
	}

	return false
}

func PartOne(input string) (int64, error) {
	calibrations, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var acc int64
	for _, c := range calibrations {
		if c.trySolve([]operation{OpAdd, OpMul}) {
			acc += c.lvalue
		}
	}

	return acc, nil
}

func PartTwo(input string) (int64, error) {
	calibrations, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var acc int64
	for _, c := range calibrations {
		if c.trySolve([]operation{OpAdd, OpMul, OpConcat}) {
			acc += c.lvalue
		}
	}

	return acc, nil
}
