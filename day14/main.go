package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sandSpawn = point{500, 0}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	lines := map[string]bool{}
	caveMap := map[int]map[int]bool{}
	// build map
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || lines[line] {
			continue
		}

		caveMap = addToMap(caveMap, line)
	}

	minX, maxX := 500, 500
	maxY := 0
	for i := range caveMap {
		if i < minX {
			minX = i
		}
		if i > maxX {
			maxX = i
		}
		for j := range caveMap[i] {
			if j > maxY {
				maxY = j
			}
		}
	}

	// part 1
	sandDropped := 0
	for {
		next := true
		caveMap, next = dropSand(caveMap, minX, maxX, maxY)
		if !next {
			break
		}
		sandDropped++
	}
	fmt.Println(sandDropped)

	//part 2
	floorY := maxY + 2
	for i := minX - floorY; i <= maxX+floorY; i++ {
		if _, ok := caveMap[i]; !ok {
			caveMap[i] = map[int]bool{}
		}
		caveMap[i][floorY] = true
	}
	for {
		next := true
		caveMap, next = dropSand(caveMap, -9999, 9999, floorY)
		sandDropped++
		if !next {
			break
		}
	}
	fmt.Println(sandDropped)

}

func addToMap(m map[int]map[int]bool, line string) map[int]map[int]bool {
	points := strings.Split(line, " -> ")

	for i := 0; i < len(points)-1; i++ {
		line := drawLine(points[i], points[i+1])
		for _, p := range line {
			if _, ok := m[p.x]; !ok {
				m[p.x] = map[int]bool{}
			}
			m[p.x][p.y] = true
		}
	}
	return m
}

func toPoint(s string) point {
	vals := strings.Split(s, ",")
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		panic(err)
	}
	return point{x, y}
}

func drawLine(point1 string, point2 string) []point {
	p1 := toPoint(point1)
	p2 := toPoint(point2)

	line := []point{
		p1,
		p2,
	}

	if p1.x < p2.x {
		for i := p1.x + 1; i < p2.x; i++ {
			line = append(line, point{i, p1.y})
		}
	}
	if p2.x < p1.x {
		for i := p2.x + 1; i < p1.x; i++ {
			line = append(line, point{i, p2.y})
		}
	}
	if p1.y < p2.y {
		for i := p1.y + 1; i < p2.y; i++ {
			line = append(line, point{p1.x, i})
		}
	}
	if p2.y < p1.y {
		for i := p2.y + 1; i < p1.y; i++ {
			line = append(line, point{p2.x, i})
		}
	}
	return line
}

func dropSand(m map[int]map[int]bool, leftBound, rightBound, bottom int) (map[int]map[int]bool, bool) {

	attemptMove := func(sand point) point {
		// sand is falling through the bottom
		if sand.y == bottom {
			return point{}
		}
		if occupied := m[sand.x][sand.y+1]; !occupied {
			return point{sand.x, sand.y + 1}
		}
		if sand.x-1 >= leftBound {
			if occupied := m[sand.x-1][sand.y+1]; !occupied {
				return point{sand.x - 1, sand.y + 1}
			}
		}
		if sand.x+1 <= rightBound {
			if occupied := m[sand.x+1][sand.y+1]; !occupied {
				return point{sand.x + 1, sand.y + 1}
			}
		}
		// sand is falling out of bounds
		if sand.x == leftBound || sand.x == rightBound {
			return point{}
		}
		return sand
	}

	sand := sandSpawn
	for {
		move := attemptMove(sand)
		// sand stopped
		if move == sand {
			m[sand.x][sand.y] = true
			// sand spawn is blocked
			if move == sandSpawn {
				return m, false
			}
			return m, true
		}
		// sand fell off
		if move.x == 0 {
			return m, false
		}
		sand = move
	}

}

type point struct {
	x, y int
}
