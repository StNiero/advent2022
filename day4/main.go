package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	contained := 0
	overlapped := 0
	for i := 1; scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), ",")
		if len(line) == 0 {
			continue
		}
		elf1 := strings.Split(line[0], "-")
		elf1min, err := strconv.Atoi(elf1[0])
		if err != nil {
			panic(err)
		}
		elf1max, err := strconv.Atoi(elf1[1])
		if err != nil {
			panic(err)
		}
		elf2 := strings.Split(line[1], "-")
		elf2min, err := strconv.Atoi(elf2[0])
		if err != nil {
			panic(err)
		}
		elf2max, err := strconv.Atoi(elf2[1])
		if err != nil {
			panic(err)
		}

		if elf1min <= elf2min {
			if elf1max >= elf2max {
				contained++
			}
			if elf1max >= elf2min {
				overlapped++
			}
			continue
		} else if elf2min <= elf1min {
			if elf2max >= elf1max {
				contained++
			}
			if elf2max >= elf1min {
				overlapped++
			}
			continue
		}
	}

	fmt.Println(contained)
	fmt.Println(overlapped)
}
