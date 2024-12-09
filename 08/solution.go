package day_08

import (
	"fmt"
	"strings"
)

type coord struct {
	y int
	x int
}

func (c coord) add(other coord) coord {
	return coord{
		y: c.y + other.y,
		x: c.x + other.x,
	}
}

func (c coord) sub(other coord) coord {
	return coord{
		y: c.y - other.y,
		x: c.x - other.x,
	}
}

type coordSet map[coord]struct{}

func (c coordSet) merge(other coordSet) coordSet {
	for k, v := range other {
		c[k] = v
	}
	return c
}

type Map struct {
	sizeY int
	sizeX int

	antennas map[string]coordSet
}

func parseInput(input string) *Map {
	lines := strings.Split(input, "\n")

	m := new(Map)
	m.sizeY = len(lines)
	m.sizeX = len(lines[0])

	m.antennas = make(map[string]coordSet)

	for yidx, line := range lines {
		for xidx, char := range line {
			if char != '.' {
				if _, ok := m.antennas[string(char)]; !ok {
					m.antennas[string(char)] = make(coordSet)
				}

				m.antennas[string(char)][coord{y: yidx, x: xidx}] = struct{}{}
			}
		}
	}

	return m
}

func (m Map) filterValidCoords(coords coordSet) coordSet {
	res := make(coordSet)

	for coord := range coords {
		if coord.y >= 0 && coord.y < m.sizeY && coord.x >= 0 && coord.x < m.sizeX {
			res[coord] = struct{}{}
		}
	}

	return res
}

func (m Map) getAntinodesForSymbolPartOne(symbol string) coordSet {
	coordSetCopy := make(coordSet)
	for k, v := range m.antennas[symbol] {
		coordSetCopy[k] = v
	}

	res := make(coordSet)

	for firstAntenna := range m.antennas[symbol] {
		delete(coordSetCopy, firstAntenna)

		for otherAntenna := range coordSetCopy {
			res[otherAntenna.sub(firstAntenna.sub(otherAntenna))] = struct{}{}
			res[firstAntenna.add(firstAntenna.sub(otherAntenna))] = struct{}{}
		}
	}

	return res
}

func (m Map) getAntinodesForSymbolPartTwo(symbol string) coordSet {
	coordSetCopy := make(coordSet)
	for k, v := range m.antennas[symbol] {
		coordSetCopy[k] = v
	}

	res := make(coordSet)

	for firstAntenna := range m.antennas[symbol] {
		delete(coordSetCopy, firstAntenna)

		for otherAntenna := range coordSetCopy {
			cur := firstAntenna
			for {
				res[cur] = struct{}{}
				cur = cur.sub(firstAntenna.sub(otherAntenna))
				if cur.y < 0 || cur.y >= m.sizeY || cur.x < 0 || cur.x >= m.sizeX {
					break
				}
			}

			cur = otherAntenna
			for {
				res[cur] = struct{}{}
				cur = cur.add(firstAntenna.sub(otherAntenna))
				if cur.y < 0 || cur.y >= m.sizeY || cur.x < 0 || cur.x >= m.sizeX {
					break
				}
			}
		}
	}

	return res
}

func (m Map) printWithAntinodes(coords coordSet) string {
	res := make([]string, 0, m.sizeY)

	for y := 0; y < m.sizeY; y++ {
		line := strings.Repeat(".", m.sizeX)

		for coord := range coords {
			if coord.y == y {
				line = line[:coord.x] + "#" + line[coord.x+1:]
			}
		}

		for antenna := range m.antennas {
			for coord := range m.antennas[antenna] {
				if coord.y == y {
					line = line[:coord.x] + antenna + line[coord.x+1:]
				}
			}
		}
		res = append(res, line)
	}

	return strings.Join(res, "\n")
}

func PartOne(input string) (int, error) {
	m := parseInput(input)

	res := make(coordSet)
	for symbol := range m.antennas {
		res = res.merge(m.getAntinodesForSymbolPartOne(symbol))
	}

	res = m.filterValidCoords(res)

	return len(res), nil
}

func PartTwo(input string) (int, error) {
	m := parseInput(input)

	res := make(coordSet)
	for symbol := range m.antennas {
		res = res.merge(m.getAntinodesForSymbolPartTwo(symbol))
	}

	res = m.filterValidCoords(res)
	fmt.Println(m.printWithAntinodes(res))

	return len(res), nil
}
