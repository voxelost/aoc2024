package day04

import (
	"strings"
)

type matrix []string

func compareVertical(substring string, lines []string, offset int) bool {
	for i, line := range lines {
		if line[offset] != substring[i] {
			return false
		}
	}
	return true
}

func (m matrix) countVertical() int {
	var acc int

	for x := range len(m[0]) {
		for y := 0; y < len(m)-3; y++ {
			if compareVertical("XMAS", m[y:y+4], x) {
				acc++
			}

			if compareVertical("SAMX", m[y:y+4], x) {
				acc++
			}
		}
	}

	return acc
}

func (m matrix) countHorizontal() int {
	if len(m) == 0 {
		return -1
	}

	var acc int

	for y := range len(m[0]) {
		acc += strings.Count(string(m[y]), "XMAS") + strings.Count(string(m[y]), "SAMX")
	}

	return acc
}

func (m matrix) countDiagonalRight(substring string, lines []string) int {
	if len(m) == 0 {
		return -1
	}

	var acc int

	for x := 0; x < len(m[0])-3; x++ {
		found := true
		for y, line := range lines {

			if line[x+y] != substring[y] {
				found = false
				break
			}
		}

		if found {
			acc++
		}
	}

	return acc
}

func (m matrix) countDiagonalLeft(substring string, lines []string) int {
	if len(m) == 0 {
		return -1
	}

	var acc int

	for x := 3; x < len(m[0]); x++ {
		found := true
		for y, line := range lines {

			if line[x-y] != substring[y] {
				found = false
				break
			}
		}

		if found {
			acc++
		}
	}

	return acc
}

func (m matrix) countDiagonal() int {
	if len(m) == 0 {
		return -1
	}

	var acc int

	for y := 0; y < len(m)-3; y++ {
		acc += m.countDiagonalRight("XMAS", m[y:y+4])
		acc += m.countDiagonalRight("SAMX", m[y:y+4])

		acc += m.countDiagonalLeft("XMAS", m[y:y+4])
		acc += m.countDiagonalLeft("SAMX", m[y:y+4])
	}

	return acc
}

func (m matrix) hasSubmatrixAt(s matrix, y, x int) bool {
	for sy, line := range s {
		for sx := range line {
			symbol := line[sx]
			if symbol != '.' && m[y+sy][x+sx] != symbol {
				return false
			}
		}
	}

	return true
}

func (m matrix) countSubmatrix(s matrix) int {
	var acc int

	for y := 0; y < len(m)-2; y++ {
		for x := 0; x < len(m[0])-2; x++ {
			if m.hasSubmatrixAt(s, y, x) {
				acc++
			}
		}
	}

	return acc
}

func PartOne(input string) (int, error) {
	m := matrix(strings.Split(input, "\n"))
	horizontal := m.countHorizontal()
	vertical := m.countVertical()
	diagonal := m.countDiagonal()

	return horizontal + vertical + diagonal, nil
}

func PartTwo(input string) (int, error) {
	matrices := []matrix{
		{
			"M.M",
			".A.",
			"S.S",
		},
		{
			"S.M",
			".A.",
			"S.M",
		},
		{
			"S.S",
			".A.",
			"M.M",
		},
		{
			"M.S",
			".A.",
			"M.S",
		},
	}

	var acc int

	m := matrix(strings.Split(input, "\n"))

	for _, s := range matrices {
		acc += m.countSubmatrix(s)
	}

	return acc, nil
}
