package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []uint64
	operation func(uint64, bool) uint64
	test      func(map[int]*Monkey, uint64)

	totalInspected int
}

var re = regexp.MustCompile("[0-9]+")

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	monkeys := map[int]*Monkey{}
	monkeyIdx := 0
	lazyLCM := uint64(1)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Monkey") {
			val := re.FindString(line)
			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			monkeyIdx = v
			if _, ok := monkeys[monkeyIdx]; !ok {
				monkeys[monkeyIdx] = &Monkey{}
			}
			continue
		}
		if strings.Contains(line, "Starting") {
			items := re.FindAllString(line, -1)
			for _, item := range items {
				i, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}
				monkeys[monkeyIdx].items = append(monkeys[monkeyIdx].items, uint64(i))
			}
			continue
		}
		if strings.Contains(line, "Operation") {
			val := re.FindString(line)

			// If we weren't given a number the input operation should be old * old
			if val == "" {
				monkeys[monkeyIdx].operation = func(worry uint64, destress bool) uint64 {
					if !destress {
						return worry * worry
					}
					return worry * worry / 3
				}
				continue
			}

			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			if strings.Contains(line, "*") {
				monkeys[monkeyIdx].operation = func(worry uint64, destress bool) uint64 {
					if !destress {
						return worry * uint64(v)
					}
					return worry * uint64(v) / 3
				}
				continue
			}
			if strings.Contains(line, "+") {
				monkeys[monkeyIdx].operation = func(worry uint64, destress bool) uint64 {
					if !destress {
						return worry + uint64(v)
					}
					return (worry + uint64(v)) / 3
				}
				continue
			}
		}
		if strings.Contains(line, "Test") {
			val := re.FindString(line)
			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}

			// Keep track of LCM of all test values for part 2
			lazyLCM *= uint64(v)

			// Identify monkey target if test result is true
			scanner.Scan()
			trueLine := scanner.Text()
			trueMonkey := re.FindString(trueLine)
			tm, err := strconv.Atoi(trueMonkey)
			if err != nil {
				panic(err)
			}
			// Identify monkey target if test result is false
			scanner.Scan()
			falseLine := scanner.Text()
			falseMonkey := re.FindString(falseLine)
			fm, err := strconv.Atoi(falseMonkey)
			if err != nil {
				panic(err)
			}

			monkeys[monkeyIdx].test = func(monkMap map[int]*Monkey, item uint64) {
				if item%uint64(v) == 0 {
					monkMap[tm].items = append(monkMap[tm].items, item)
				} else {
					monkMap[fm].items = append(monkMap[fm].items, item)
				}
			}
		}
	}

	// Make a copy of the input result for part 2
	monkeysCopy := map[int]*Monkey{}
	for k, v := range monkeys {
		m := *v
		monkeysCopy[k] = &m
	}

	for round := 0; round < 20; round++ {
		for mIdx := 0; mIdx < len(monkeys); mIdx++ {
			monkey := monkeys[mIdx]
			for i := 0; i < len(monkey.items); i++ {
				item := monkey.operation(monkey.items[i], true)
				monkey.totalInspected++
				monkey.test(monkeys, item)
			}
			monkey.items = []uint64{}
		}
	}

	inspections := []int{}
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.totalInspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	fmt.Println(inspections[0] * inspections[1])

	// part 2
	for round := 0; round < 10000; round++ {
		for mIdx := 0; mIdx < len(monkeysCopy); mIdx++ {
			monkey := monkeysCopy[mIdx]
			for i := 0; i < len(monkey.items); i++ {
				// inspect item to get new worry value, "manage our stress" along the way
				item := monkey.operation(monkey.items[i]%lazyLCM, false)
				monkey.totalInspected++
				monkey.test(monkeysCopy, item)
			}
			monkey.items = []uint64{}
		}
	}

	inspections2 := []int{}
	for _, monkey := range monkeysCopy {
		inspections2 = append(inspections2, monkey.totalInspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections2)))
	fmt.Println(inspections2[0] * inspections2[1])

}
