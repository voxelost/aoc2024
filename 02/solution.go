package day_02

import (
	"math"
	"strconv"
	"strings"
)

func checkReport(reportInts []int) bool {
	if len(reportInts) == 1 {
		return true
	}

	var assumeIncreasing bool
	for i, report := range reportInts {
		if i > 0 {
			if reportInts[i-1] > report {
				break
			} else if reportInts[i-1] < report {
				assumeIncreasing = true
				break
			}
		}
	}

	safe := true
	for i, report := range reportInts {
		if i > 0 {
			if assumeIncreasing && reportInts[i-1] > report || !assumeIncreasing && reportInts[i-1] < report {
				safe = false
				break
			}

			diff := int(math.Abs(float64(report - reportInts[i-1])))
			if diff < 1 || diff > 3 {
				safe = false
				break
			}
		}
	}
	return safe
}

func PartOne(input string) (int, error) {
	lines := strings.Split(input, "\n")

	var safeReports int
	for _, line := range lines {
		reports := strings.Split(line, " ")
		var reportInts []int
		for _, report := range reports {
			reportInt, err := strconv.Atoi(report)
			if err != nil {
				return 0, err
			}
			reportInts = append(reportInts, reportInt)
		}

		if checkReport(reportInts) {
			safeReports++
		}
	}

	return safeReports, nil
}

func getReportsWithTolerance(reportInts []int) [][]int {
	out := [][]int{reportInts}

	l := len(reportInts)

	for i := 0; i < l; i++ {
		singleReport := make([]int, l)

		copy(singleReport, reportInts)
		singleReport = append(singleReport[:i], singleReport[i+1:]...)
		out = append(out, singleReport)
	}

	return out
}

func PartTwo(input string) (int, error) {
	lines := strings.Split(input, "\n")
	var safeReports int

	for _, line := range lines {
		reports := strings.Split(line, " ")
		var reportInts []int
		for _, report := range reports {
			reportInt, err := strconv.Atoi(report)
			if err != nil {
				return 0, err
			}
			reportInts = append(reportInts, reportInt)
		}

		reportsWithTolarance := getReportsWithTolerance(reportInts)

		var safe bool
		for _, reportInts := range reportsWithTolarance {
			if checkReport(reportInts) {
				safe = true
				break
			}
		}

		if safe {
			safeReports++
		}
	}

	return safeReports, nil
}
