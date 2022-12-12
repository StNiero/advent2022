package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	heightMap := [][]int{}
	var start point
	var end point
	var starts []point
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var lineValues []int
		for _, v := range line {
			if v == 'S' {
				start = point{len(lineValues), len(heightMap)}
				v = 'a'
			}
			if v == 'E' {
				end = point{len(lineValues), len(heightMap)}
				v = 'z'
			}
			if v == 'a' {
				starts = append(starts, point{len(lineValues), len(heightMap)})
			}
			lineValues = append(lineValues, int(v-97))
		}
		heightMap = append(heightMap, lineValues)
	}

	fmt.Println(calcShortest(heightMap, start, end))

	shortestDistance := 999999999
	for _, s := range starts {
		if dist := calcShortest(heightMap, s, end); dist < shortestDistance {
			shortestDistance = dist
		}
	}
	fmt.Println(shortestDistance)
}

func calcShortest(heightMap [][]int, start point, target point) int {
	visited := map[point]int{}
	queue := []point{start}

	for len(queue) > 0 {
		cur := queue[0]
		if cur == target {
			return visited[cur]
		}
		queue = queue[1:]

		// If theres a point to the left and it's 1 or less higher than the current point, explore it
		// but only if we haven't been there already
		if cur.x > 0 && heightMap[cur.y][cur.x-1] <= heightMap[cur.y][cur.x]+1 {
			next := point{cur.x - 1, cur.y}
			if _, ok := visited[next]; !ok {
				queue = append(queue, next)
				visited[next] = visited[cur] + 1
			}
		}
		// If theres a point to the right and it's 1 or less higher than the current point, explore it
		// but only if we haven't been there already
		if cur.x < len(heightMap[cur.y])-1 && heightMap[cur.y][cur.x+1] <= heightMap[cur.y][cur.x]+1 {
			next := point{cur.x + 1, cur.y}
			if _, ok := visited[next]; !ok {
				queue = append(queue, next)
				visited[next] = visited[cur] + 1
			}
		}
		// If theres a point above and it's 1 or less higher than the current point, explore it
		// but only if we haven't been there already
		if cur.y > 0 && heightMap[cur.y-1][cur.x] <= heightMap[cur.y][cur.x]+1 {
			next := point{cur.x, cur.y - 1}
			if _, ok := visited[next]; !ok {
				queue = append(queue, next)
				visited[next] = visited[cur] + 1
			}
		}
		// If theres a point below and it's 1 or less higher than the current point, explore it
		// but only if we haven't been there already
		if cur.y < len(heightMap)-1 && heightMap[cur.y+1][cur.x] <= heightMap[cur.y][cur.x]+1 {
			next := point{cur.x, cur.y + 1}
			if _, ok := visited[next]; !ok {
				queue = append(queue, next)
				visited[next] = visited[cur] + 1
			}
		}
	}

	// No path
	return 999999999
}
