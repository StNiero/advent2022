package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	// parse drawing lines
	var rawDrawing [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rawDrawing = append(rawDrawing, parseDrawingLine(line))
	}

	// work backwards to convert into stacks, skip the column numbers
	stacks := make(map[int]*Stack)
	for l := len(rawDrawing) - 2; l >= 0; l-- {
		for i, val := range rawDrawing[l] {
			if val == " " {
				continue
			}
			if _, ok := stacks[i+1]; !ok {
				stacks[i+1] = &Stack{}
			}
			stacks[i+1].Push(val)
		}
	}

	// copy for CrateMover 9001 procedure (part2)
	stacks2 := make(map[int]*Stack)
	for k := range stacks {
		if _, ok := stacks2[k]; !ok {
			stacks2[k] = &Stack{}
		}
		for _, v := range *stacks[k] {
			stacks2[k].Push(v)
		}
	}

	// perform procedure lines
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		procedure := parseProcedureLine(line)
		move := procedure[0]
		from := procedure[1]
		to := procedure[2]

		// Part 1
		for i := 0; i < move; i++ {
			val := stacks[from].Pop()
			stacks[to].Push(val)
		}

		// Part2
		vals := stacks2[from].BulkPop(move)
		stacks2[to].MultiPush(vals)
	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i].Pop())
	}
	fmt.Println()
	for i := 1; i <= len(stacks2); i++ {
		fmt.Print(stacks2[i].Pop())
	}
}

func parseDrawingLine(line string) []string {
	rawLine := []string{}
	for i := 0; i < len(line); i++ {
		if i%4 == 1 {
			rawLine = append(rawLine, string(line[i]))
		}
	}
	return rawLine
}

func parseProcedureLine(line string) []int {
	// trim out non-digits
	re := regexp.MustCompile("[0-9]+")
	stringProc := re.FindAllString(line, -1)

	proc := []int{}
	for _, s := range stringProc {
		val, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		proc = append(proc, val)
	}
	return proc
}

type Stack []string

func (s *Stack) Push(val string) {
	*s = append(*s, val)
}

func (s *Stack) Pop() string {
	length := len(*s)
	if length == 0 {
		return ""
	}
	val := (*s)[length-1]
	*s = (*s)[:length-1]
	return val
}

// MultiPush pushes a slice of strings onto the stack in the order given
func (s *Stack) MultiPush(vals []string) {
	for _, v := range vals {
		s.Push(v)
	}
}

// BulkPop returns a slice of values popped from the stack, retaining their order
// i.e. Given [5 4 3 2], BulkPop(2) returns [3 2]
// This is NOT the same as calling Pop() twice
func (s *Stack) BulkPop(num int) []string {
	length := len(*s)
	if length == 0 {
		return []string{}
	}
	// can ignore num > length, input doesn't have that problem
	vals := (*s)[length-num:]
	*s = (*s)[:length-num]
	return vals
}
