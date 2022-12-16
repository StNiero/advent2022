package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const targetY = 2000000
const maxCoord = 4000000

// const targetY = 10
// const maxCoord = 20

type minMax struct {
	min, max int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	tunnelMap1 := map[point]rune{}
	intervals := map[int][]minMax{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		tunnelMap1 = addToMap(tunnelMap1, line)
		intervals = addToIntervals(intervals, line)
	}

	noBeacon := 0
	for coord := range tunnelMap1 {
		if coord.y == targetY && tunnelMap1[point{coord.x, coord.y}] == '#' {
			noBeacon++
		}

	}

	fmt.Println(noBeacon)

	for row := 0; row <= maxCoord; row++ {
		i := merge(intervals[row])
		if len(i) > 1 {
			fmt.Println(tuningFrequency(point{i[0].max + 1, row}))
		}
	}
}

func addToMap(m map[point]rune, line string) map[point]rune {
	re := regexp.MustCompile(`-*\d+`)
	coords := re.FindAllString(line, -1)
	sensor := toPoint(coords[0], coords[1])
	beacon := toPoint(coords[2], coords[3])

	m[sensor] = 'S'
	m[beacon] = 'B'

	distance := abs(beacon.x-sensor.x) + abs(beacon.y-sensor.y)
	// don't waste time if we're not drawing near the y coord in question
	if abs(targetY-sensor.y) > distance {
		return m
	}

	// map out the points that can't contain a beacon
	for i := 0; i < distance; i++ {
		// but only if they fall on the target row
		if sensor.y+i == targetY {
			for _, p := range drawLine(point{sensor.x + distance - i, sensor.y + i}, point{sensor.x - distance + i, sensor.y + i}) {
				if _, ok := m[p]; !ok {
					m[p] = '#'
				}
			}
		}
		// again only if they fall on the target row
		if sensor.y-i == targetY {
			for _, p := range drawLine(point{sensor.x + distance - i, sensor.y - i}, point{sensor.x - distance + i, sensor.y - i}) {
				if _, ok := m[p]; !ok {
					m[p] = '#'
				}
			}
		}
	}

	return m
}

func addToIntervals(ints map[int][]minMax, line string) map[int][]minMax {
	re := regexp.MustCompile(`-*\d+`)
	coords := re.FindAllString(line, -1)
	sensor := toPoint(coords[0], coords[1])
	beacon := toPoint(coords[2], coords[3])

	distance := abs(beacon.x-sensor.x) + abs(beacon.y-sensor.y)

	for i := 0; i < distance; i++ {
		if sensor.y+i <= maxCoord {
			min := sensor.x - distance + i
			if min < 0 {
				min = 0
			}
			max := sensor.x + distance - i
			if max > maxCoord {
				max = maxCoord
			}
			if len(ints[sensor.y+i]) == 0 {
				ints[sensor.y+i] = []minMax{}
			}
			ints[sensor.y+i] = append(ints[sensor.y+i], minMax{
				min: min,
				max: max,
			})
		}
		if sensor.y-i >= 0 {
			min := sensor.x - distance + i
			if min < 0 {
				min = 0
			}
			max := sensor.x + distance - i
			if max > maxCoord {
				max = maxCoord
			}
			ints[sensor.y-i] = append(ints[sensor.y-i], minMax{
				min: min,
				max: max,
			})
		}
	}

	return ints
}

func merge(intervals []minMax) []minMax {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].min < intervals[j].min
	})

	new := []minMax{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i].max < new[len(new)-1].max {
			continue
		}
		if intervals[i].min <= new[len(new)-1].max+1 {
			new[len(new)-1].max = intervals[i].max
			continue
		}
		new = append(new, intervals[i])
	}
	return new
}

func toPoint(xStr, yStr string) point {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		panic(err)
	}
	return point{x, y}
}

func drawLine(p1, p2 point) []point {

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

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

func tuningFrequency(p point) int {
	return p.x*4000000 + p.y
}

type point struct {
	x, y int
}
