package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	rope := make([]struct {
		x int
		y int
	}, 10)
	tailVisited := map[int]map[int]bool{}
	tail9Visited := map[int]map[int]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		row := strings.Split(line, " ")
		dir := row[0]
		steps, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}
		for i := 0; i < steps; i++ {
			switch dir {
			case "L":
				rope[0].x--
			case "R":
				rope[0].x++
			case "U":
				rope[0].y++
			case "D":
				rope[0].y--
			}

			for i := 1; i < len(rope); i++ {
				parentX := &rope[i-1].x
				parentY := &rope[i-1].y
				childX := &rope[i].x
				childY := &rope[i].y
				touching := math.Abs(float64(*parentX-*childX)) <= 1 && math.Abs(float64(*parentY-*childY)) <= 1
				// If not touching and not sharing a row/column move diagonally
				if !touching && *parentX != *childX && *parentY != *childY {
					*childX += (*parentX - *childX) / int(math.Abs(float64(*parentX-*childX)))
					*childY += (*parentY - *childY) / int(math.Abs(float64(*parentY-*childY)))
					continue
				}

				// Otherwise move one direction
				if math.Abs(float64(*parentX-*childX)) > 1 {
					*childX += (*parentX - *childX) / int(math.Abs(float64(*parentX-*childX)))
				}
				if math.Abs(float64(*parentY-*childY)) > 1 {
					*childY += (*parentY - *childY) / int(math.Abs(float64(*parentY-*childY)))
				}
			}

			// Part1
			if tailVisited[rope[1].x] == nil {
				tailVisited[rope[1].x] = map[int]bool{}
			}
			tailVisited[rope[1].x][rope[1].y] = true

			// Part2
			if tail9Visited[rope[9].x] == nil {
				tail9Visited[rope[9].x] = map[int]bool{}
			}
			tail9Visited[rope[9].x][rope[9].y] = true
		}
	}
	total := 0
	for _, points := range tailVisited {
		total += len(points)
	}
	total2 := 0
	for _, points := range tail9Visited {
		total2 += len(points)
	}

	fmt.Println(total, total2)
}
