package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	var elves []int
	for i := 0; scanner.Scan(); {
		line := scanner.Text()
		if line == "" {
			i++
			continue
		}
		cal, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if len(elves) == i {
			elves = append(elves, 0)
		}
		elves[i] += cal
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Println(elves[0])
	fmt.Println(elves[0] + elves[1] + elves[2])
}
