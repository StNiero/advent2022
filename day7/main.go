package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	diskSpace   = 70000000
	spaceNeeded = 30000000
)

var root = &Node{
	name:     "/",
	children: map[string]*Node{},
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	var pwd *Node
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "$ ls") {
			continue
		}

		// We're changing the present working directory
		if strings.HasPrefix(line, "$ cd") {
			pwd = changeDir(pwd, line[5:])
			continue
		}

		// Otherwise we're printing a directory listing
		if strings.HasPrefix(line, "dir") {
			dirName := line[4:]
			if _, ok := pwd.children[dirName]; !ok {
				pwd.children[dirName] = &Node{
					name:     dirName,
					parent:   pwd,
					children: map[string]*Node{},
				}
			}
			continue
		}

		fileLine := strings.Split(line, " ")
		fileSize, err := strconv.Atoi(fileLine[0])
		if err != nil {
			panic(err)
		}
		fileName := fileLine[1]
		if _, ok := pwd.children[fileName]; !ok {
			pwd.children[fileName] = &Node{
				name:     fileName,
				size:     fileSize,
				parent:   pwd,
				children: map[string]*Node{},
			}
		}
	}

	targetTotal := 0
	dirSizes := map[string]int{}
	var getSumChildren func(n *Node) int
	getSumChildren = func(n *Node) int {
		sum := 0
		for i := range n.children {
			c := n.children[i]
			if len(c.children) > 0 {
				sum += getSumChildren(c)
				continue
			}
			sum += c.size
		}
		if sum <= 100000 {
			targetTotal += sum
		}
		dirSizes[n.name] = sum
		return sum
	}

	spaceUsed := getSumChildren(root)
	fmt.Println(targetTotal)

	freeSpace := diskSpace - spaceUsed
	spaceToClear := spaceNeeded - freeSpace

	sizeOfDir := spaceUsed
	for _, s := range dirSizes {
		if s > spaceToClear && s < sizeOfDir {
			sizeOfDir = s
		}
	}
	fmt.Println(sizeOfDir)
}

func changeDir(n *Node, dirName string) *Node {
	if dirName == "/" {
		return root
	}
	if dirName == ".." {
		return n.parent
	}

	for _, c := range n.children {
		if c.name == dirName {
			return c
		}
	}
	child := &Node{
		name:   dirName,
		parent: n,
	}
	n.children[child.name] = child
	return child
}

type Node struct {
	name     string
	size     int
	children map[string]*Node
	parent   *Node
}
