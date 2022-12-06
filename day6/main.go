package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		last4 := []rune{}
		last14 := []rune{}
		startPacket := 0
		startMsg := 0
		for i, r := range line {
			if startPacket == 0 {
				last4 = append(last4, r)
				if len(last4) > 4 {
					last4 = last4[1:]
				} else {
					continue
				}

				dupIn4 := false
				for i := 1; i < len(last4); i++ {
					if strings.Count(string(last4), string(last4[i])) > 1 {
						dupIn4 = true
						break
					}
				}
				if !dupIn4 {
					startPacket = i + 1
				}
			}

			last14 = append(last14, r)
			if len(last14) > 14 {
				last14 = last14[1:]
			} else {
				continue
			}

			dupIn14 := false
			for i := 1; i < len(last14); i++ {
				if strings.Count(string(last14), string(last14[i])) > 1 {
					dupIn14 = true
					break
				}
			}
			if !dupIn14 {
				startMsg = i + 1
				break
			}
		}
		fmt.Println(startPacket, startMsg)
	}
}
