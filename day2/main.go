package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	opRock    = "A"
	opPaper   = "B"
	opScissor = "C"
	rock      = "X"
	paper     = "Y"
	scissor   = "Z"

	lose = "X"
	draw = "Y"
	win  = "Z"
)

var playScore = map[string]int{
	rock:    1,
	paper:   2,
	scissor: 3,
}

var resultScore = map[string]int{
	win:  6,
	draw: 3,
	lose: 0,
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	var score1, score2 int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		round := strings.Split(line, " ")
		score1 += playStrat1(round[0], round[1])
		score2 += playStrat2(round[0], round[1])
	}

	fmt.Println(score1)
	fmt.Println(score2)
}

func playStrat1(opPlay, elfPlay string) int {
	score := playScore[elfPlay]

	switch elfPlay {
	case rock:
		switch opPlay {
		case opRock:
			score += resultScore[draw]
		case opPaper:
			score += resultScore[lose]
		case opScissor:
			score += resultScore[win]
		}
	case paper:
		switch opPlay {
		case opRock:
			score += resultScore[win]
		case opPaper:
			score += resultScore[draw]
		case opScissor:
			score += resultScore[lose]
		}
	case scissor:
		switch opPlay {
		case opRock:
			score += resultScore[lose]
		case opPaper:
			score += resultScore[win]
		case opScissor:
			score += resultScore[draw]
		}
	}
	return score
}

func playStrat2(opPlay, result string) int {
	score := resultScore[result]

	switch result {
	case win:
		switch opPlay {
		case opRock:
			score += playScore[paper]
		case opPaper:
			score += playScore[scissor]
		case opScissor:
			score += playScore[rock]
		}
	case lose:
		switch opPlay {
		case opRock:
			score += playScore[scissor]
		case opPaper:
			score += playScore[rock]
		case opScissor:
			score += playScore[paper]
		}
	case draw:
		switch opPlay {
		case opRock:
			score += playScore[rock]
		case opPaper:
			score += playScore[paper]
		case opScissor:
			score += playScore[scissor]
		}
	}

	return score
}
