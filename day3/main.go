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

	alphabet := " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	sum := 0
	badgeSum := 0
	group := []string{}
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			continue
		}
		length := len(line)
		ruck1 := line[0 : length/2]
		ruck2 := line[length/2 : length]

		c := intersectChars(ruck1, ruck2)
		sum += strings.Index(alphabet, c)

		group = append(group, line)
		if i%3 == 0 {
			badgeSum += strings.Index(alphabet, intersectChars(intersectAllChars(group[0], group[1]), group[2]))
			group = []string{}
		}
	}

	fmt.Println(sum)
	fmt.Println(badgeSum)
}

// intersectChars takes 2 strings and returns the first shared character between them as a string
// it returns "0" if there are no shared characters
func intersectChars(a, b string) string {
	for _, x := range strings.Split(a, "") {
		if strings.Contains(b, x) {
			return x
		}
	}
	return "0"
}

// intersectAllChars takes 2 strings and returns all shared characters between them as a string
func intersectAllChars(a, b string) string {
	intersect := ""
	for _, x := range strings.Split(a, "") {
		if strings.Contains(b, x) {
			intersect += x
		}
	}
	return intersect
}
