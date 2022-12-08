package main

import (
	"bufio"
	"fmt"
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

	trees := [][]int{}
	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		line := scanner.Text()
		if line == "" {
			continue
		}

		row := strings.Split(line, "")
		treeRow := []int{}
		for _, v := range row {
			height, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			treeRow = append(treeRow, height)
		}
		trees = append(trees, treeRow)
	}

	numVisible := 0
	highestScore := 0
	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[row]); col++ {

			var isVisible = func(row, col int) bool {
				if row == 0 || row == len(trees)-1 || col == 0 || col == len(trees[row])-1 {
					return true
				}
				treeHeight := trees[row][col]
				visFromTop, visFromBottom := true, true
				for i := range trees {
					if i < row && trees[i][col] >= treeHeight {
						visFromTop = false
					}
					if i > row && trees[i][col] >= treeHeight {
						visFromBottom = false
					}
				}
				if visFromTop || visFromBottom {
					return true
				}
				visFromLeft, visFromRight := true, true
				for i := range trees[row] {
					if i < col && trees[row][i] >= treeHeight {
						visFromLeft = false

					}
					if i > col && trees[row][i] >= treeHeight {
						visFromRight = false
					}
				}
				if visFromRight || visFromLeft {
					return true
				}
				return false
			}
			if isVisible(row, col) {
				numVisible++
			}

			var scenicScore = func(row, col int) int {
				treeHeight := trees[row][col]

				viewUp := 0
				for i := row - 1; i >= 0; i-- {
					viewUp++
					if trees[i][col] < treeHeight {
						continue
					}
					break
				}
				viewDown := 0
				for i := row + 1; i < len(trees); i++ {
					viewDown++
					if trees[i][col] < treeHeight {
						continue
					}
					break
				}
				viewLeft := 0
				for i := col - 1; i >= 0; i-- {
					viewLeft++
					if trees[row][i] < treeHeight {
						continue
					}
					break
				}
				viewRight := 0
				for i := col + 1; i < len(trees[row]); i++ {
					viewRight++
					if trees[row][i] < treeHeight {
						continue
					}
					break
				}
				return viewUp * viewDown * viewLeft * viewRight
			}
			if ss := scenicScore(row, col); ss > highestScore {
				highestScore = ss
			}
		}
	}
	fmt.Println(numVisible)
	fmt.Println(highestScore)
}
