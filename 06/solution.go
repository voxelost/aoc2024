package day_06

import (
	"fmt"
	"strings"
)

type coord struct {
	y, x int
}

func (c coord) Is(other coord) bool {
	return c.y == other.y && c.x == other.x
}

func (c coord) Moved(y, x int) coord {
	return coord{c.y + y, c.x + x}
}

type direction string

var (
	Up    direction = "^"
	Right direction = ">"
	Down  direction = "v"
	Left  direction = "<"
)

type visit struct {
	coord
	Dir direction
}

func (v visit) Is(v2 visit) bool {
	return v.coord.Is(v2.coord) && v.Dir == v2.Dir
}

type Map struct {
	Guard    visit
	Fields   [][]byte
	Obstacle coord
	Visits   map[visit]struct{}
}

func (m Map) DebugPrint() {
	var debugMap strings.Builder
	for i, line := range m.Fields {
		var printLine strings.Builder
		for j, c := range line {
			if m.Guard.coord.Is(coord{i, j}) {
				printLine.WriteString(string(m.Guard.Dir))
			} else if m.Obstacle.Is(coord{i, j}) {
				printLine.WriteString("O")
			} else if (m.VisitExists(visit{coord{i, j}, Up}) || m.VisitExists(visit{coord{i, j}, Down})) &&
				(m.VisitExists(visit{coord{i, j}, Right}) || m.VisitExists(visit{coord{i, j}, Left})) {
				printLine.WriteString("+")
			} else if m.VisitExists(visit{coord{i, j}, Up}) || m.VisitExists(visit{coord{i, j}, Down}) {
				printLine.WriteString("|")
			} else if m.VisitExists(visit{coord{i, j}, Right}) || m.VisitExists(visit{coord{i, j}, Left}) {
				printLine.WriteString("-")
			} else {
				if c == '.' {
					printLine.WriteString(" ")
				} else {
					printLine.WriteByte(c)
				}
			}
		}
		debugMap.WriteString(printLine.String())
		debugMap.WriteString("\n")
	}
	fmt.Println(debugMap.String())
}

func (m *Map) Visit(v visit) {
	m.Visits[v] = struct{}{}
	m.Guard = v
}

func (m *Map) CheckCollision(v visit) bool {
	switch v.Dir {
	case Up:
		if v.y == 0 {
			return false
		}
		if m.Fields[v.y-1][v.x] == '#' || v.Moved(-1, 0).Is(m.Obstacle) {
			return true
		}
	case Right:
		if v.x == len(m.Fields[0])-1 {
			return false
		}
		if m.Fields[v.y][v.x+1] == '#' || v.Moved(0, 1).Is(m.Obstacle) {
			return true
		}
	case Down:
		if v.y == len(m.Fields)-1 {
			return false
		}
		if m.Fields[v.y+1][v.x] == '#' || v.Moved(1, 0).Is(m.Obstacle) {
			return true
		}
	case Left:
		if v.x == 0 {
			return false
		}
		if m.Fields[v.y][v.x-1] == '#' || v.Moved(0, -1).Is(m.Obstacle) {
			return true
		}
	}

	return false
}

func (m *Map) NextVisit(v visit) (nextVisit visit, wouldLeaveMap bool) {
	nextVisit = v
	switch nextVisit.Dir {
	case Up:
		if m.CheckCollision(v) {
			nextVisit.Dir = Right
			return nextVisit, false
		}
		nextVisit.coord = nextVisit.Moved(-1, 0)
	case Right:
		if m.CheckCollision(v) {
			nextVisit.Dir = Down
			return nextVisit, false
		}
		nextVisit.coord = nextVisit.Moved(0, 1)
	case Down:
		if m.CheckCollision(v) {
			nextVisit.Dir = Left
			return nextVisit, false
		}
		nextVisit.coord = nextVisit.Moved(1, 0)
	case Left:
		if m.CheckCollision(v) {
			nextVisit.Dir = Up
			return nextVisit, false
		}
		nextVisit.coord = nextVisit.Moved(0, -1)
	}
	if nextVisit.y < 0 || nextVisit.y >= len(m.Fields) || nextVisit.x < 0 || nextVisit.x >= len(m.Fields[0]) {
		return nextVisit, true
	}
	return nextVisit, false
}

func (m *Map) VisitExists(v visit) bool {
	_, ok := m.Visits[v]
	return ok
}

func (m *Map) PlaceObstacle(c coord) {
	m.Obstacle = c
}

func (m Map) DeepCopy() *Map {
	fields := make([][]byte, len(m.Fields))
	for i, line := range m.Fields {
		fields[i] = make([]byte, len(line))
		copy(fields[i], line)
	}
	visits := make(map[visit]struct{})
	for k := range m.Visits {
		visits[k] = struct{}{}
	}

	return &Map{
		Guard:    m.Guard,
		Fields:   fields,
		Obstacle: m.Obstacle,
		Visits:   visits,
	}
}

func (m *Map) RunUntilCycle() (wouldCycle bool) {
	for {
		// m.DebugPrint()
		nextVisit, wouldLeaveMap := m.NextVisit(m.Guard)
		if wouldLeaveMap {
			// m.DebugPrint()
			return false
		}

		if m.VisitExists(nextVisit) {
			// m.DebugPrint()
			return true
		}
		m.Visit(nextVisit)
	}
}

func parseMap(input string) *Map {
	lines := strings.Split(input, "\n")
	m := &Map{
		Fields:   make([][]byte, len(lines)),
		Visits:   make(map[visit]struct{}),
		Obstacle: coord{-1, -1},
	}

	for i, line := range lines {
		if idx := strings.Index(line, string(Up)); idx > -1 {
			m.Guard = visit{coord{i, idx}, Up}
			m.Visits[m.Guard] = struct{}{}
			line = strings.ReplaceAll(line, string(m.Guard.Dir), ".")
		} else if idx := strings.Index(line, string(Down)); idx > -1 {
			m.Guard = visit{coord{i, idx}, Down}
			m.Visits[m.Guard] = struct{}{}
			line = strings.ReplaceAll(line, string(m.Guard.Dir), ".")
		} else if idx := strings.Index(line, string(Left)); idx > -1 {
			m.Guard = visit{coord{i, idx}, Left}
			m.Visits[m.Guard] = struct{}{}
			line = strings.ReplaceAll(line, string(m.Guard.Dir), ".")
		} else if idx := strings.Index(line, string(Right)); idx > -1 {
			m.Guard = visit{coord{i, idx}, Right}
			m.Visits[m.Guard] = struct{}{}
			line = strings.ReplaceAll(line, string(m.Guard.Dir), ".")
		}
		m.Fields[i] = []byte(line)
	}
	return m
}

func PartOne(input string) (int, error) {
	m := parseMap(input)

	for {
		// m.DebugPrint()
		nextVisit, didLeaveMap := m.NextVisit(m.Guard)
		if didLeaveMap {
			break
		}
		m.Visit(nextVisit)
	}

	var visitedCoords []coord
	for v := range m.Visits {
		shouldAdd := true
		for _, c := range visitedCoords {
			if c.Is(v.coord) {
				shouldAdd = false
				break
			}
		}
		if shouldAdd {
			visitedCoords = append(visitedCoords, v.coord)
		}
	}

	return len(visitedCoords), nil
}

func PartTwo(input string) (int, error) {
	m := parseMap(input)
	initialGuardLocation := m.Guard

	validObstacles := make(map[coord]struct{})
	lastObstaclePlacing := m.Obstacle
	for {
		mc := m.DeepCopy()
		nextVisit, wouldLeaveMap := mc.NextVisit(mc.Guard)
		if wouldLeaveMap {
			break
		}
		if lastObstaclePlacing.Is(nextVisit.coord) {
			// if only rotating, next placing would be on the same coord
			m.Visit(nextVisit)
			continue
		}

		mc.PlaceObstacle(nextVisit.coord)
		lastObstaclePlacing = mc.Obstacle

		if nextVisit.coord.Is(initialGuardLocation.coord) {
			continue
		}

		obstacleOnAlreadyVisited := false
		// obstacle couldn't be placed on coord we already visited
		for v := range mc.Visits {
			if v.coord.Is(lastObstaclePlacing) {
				m.Visit(nextVisit)
				obstacleOnAlreadyVisited = true
				break
			}
		}
		if obstacleOnAlreadyVisited {
			continue
		}

		wouldCycle := mc.RunUntilCycle()
		if wouldCycle {
			// mc.DebugPrint()
			validObstacles[mc.Obstacle] = struct{}{}
		}
		m.Visit(nextVisit)
	}

	return len(validObstacles), nil
}
