package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	cycle int
	x     int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	state := State{
		cycle: 1,
		x:     1,
	}
	signalStrength := 0
	crt := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		row := strings.Split(line, " ")
		cmd := row[0]

		// noop or processing 1 cycle for addx
		drawPixel(&crt, state.x)
		state.cycle++
		if state.cycle%40 == 20 {
			signalStrength += state.x * state.cycle
		}

		if cmd == "addx" {
			val, err := strconv.Atoi(row[1])
			if err != nil {
				panic(err)
			}
			// addx on the second cycle
			drawPixel(&crt, state.x)
			state = State{
				x:     state.x + val,
				cycle: state.cycle + 1,
			}
			if state.cycle%40 == 20 {
				signalStrength += state.x * state.cycle
			}
		}
	}

	fmt.Println(signalStrength)
	fmt.Print(crt)
}

func drawPixel(crt *string, x int) {
	pixel := len(strings.ReplaceAll(*crt, "\n", "")) % 40
	if pixel == 0 {
		*crt += "\n"
	}
	if pixel >= x-1 && pixel <= x+1 {
		*crt += "#"
	} else {
		*crt += "."
	}
}
