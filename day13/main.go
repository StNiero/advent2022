package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	sum := 0
	packets := []string{}
	for i := 1; scanner.Scan(); {
		left := scanner.Text()
		if left == "" {
			continue
		}
		scanner.Scan()
		right := scanner.Text()

		if v := compare(left, right); v < 0 {
			sum += i
		}

		packets = append(packets, left, right)
		i++
	}
	fmt.Println(sum)

	div1 := "[[2]]"
	div2 := "[[6]]"
	packets = append(packets, div1, div2)
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})
	part2 := 1
	for i := range packets {
		if packets[i] == div1 || packets[i] == div2 {
			part2 *= i + 1
		}
	}
	fmt.Println(part2)
}

func compare(left, right string) int {
	if left == right {
		return 0
	}

	// They are both numbers, we can compare
	if !strings.Contains(left, "[") && !strings.Contains(right, "[") {
		return compareAsInts(left, right)
	}

	leftVals := parsePacket(left)
	rightVals := parsePacket(right)

	for i := 0; i < len(leftVals); i++ {
		// Right side ran out of items, so inputs are not in the right order
		if len(rightVals) < i+1 {
			return 1
		}
		l := leftVals[i]
		r := rightVals[i]

		if l == "" {
			return -1 // Left side ran out of items, so inputs are in the right order
		}
		if r == "" {
			return 1 // Right side ran out of items, so inputs are not in the right order
		}

		if !strings.Contains(l, "[") && !strings.Contains(r, "[") {
			if v := compare(l, r); v != 0 {
				return v
			}
			continue
		}

		if !strings.Contains(l, "[") {
			if v := compare("["+l+"]", r); v != 0 {
				return v
			}
			continue
		}
		if !strings.Contains(r, "[") {
			if v := compare(l, "["+r+"]"); v != 0 {
				return v
			}
			continue
		}
		if v := compare(l, r); v != 0 {
			return v
		}

	}

	// Left side ran out of items, so inputs are in the right order
	return -1
}

func parsePacket(packet string) []string {
	packet = strings.TrimSuffix(strings.TrimPrefix(packet, "["), "]")
	if packet == "" {
		return []string{""}
	}

	var values []string
	var value string
	brackets := 0
	for _, char := range strings.Split(packet, "") {
		value += char
		if char == "[" {
			brackets++
		}
		if char == "]" {
			brackets--
		}
		if char == "," {
			if brackets == 0 {
				values = append(values, strings.TrimSuffix(value, ","))
				value = ""
			}
		}
	}
	if value != "" {
		values = append(values, value)
	}

	return values
}

func compareAsInts(l, r string) int {
	lInt, err := strconv.Atoi(l)
	if err != nil {
		panic(err)
	}
	rInt, err := strconv.Atoi(r)
	if err != nil {
		panic(err)
	}

	if lInt < rInt {
		return -1
	}
	if lInt > rInt {
		return 1
	}
	return 0
}
