package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"./pkg/tree"
)

func main() {
	t := tree.NewTree()
	_ = read("input.txt", t)
	fmt.Printf("%v\n", len(t.EndNodes))
	fmt.Printf("%v\n", len(t.Nodes))
	answer := t.CountOrbits()
	println(answer)

	// t := tree.NewTree()
	// _ = read("test.txt", t)
	// answer := t.CountOrbits()
	// println(answer)
	// println(t.PrettyPrint())
	println(t.Distance("YOU", "SAN"))
}

func read(input string, t *tree.Tree) error {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ")")
		p := s[0]
		n := s[1]
		// println(p, n)

		var node *tree.Node
		var parent *tree.Node

		if t.Nodes[p] != nil {
			parent = t.Nodes[p]
		} else {
			parent = &tree.Node{Name: p, Children: make([]*tree.Node, 0)}
		}
		t.Nodes[p] = parent

		node = &tree.Node{Name: n, Children: make([]*tree.Node, 0)}
		if t.Nodes[n] != nil {
			node = t.Nodes[n]
		}
		node.Parent = parent
		parent.Children = append(parent.Children, node)
		t.Nodes[n] = node
		// fmt.Printf("%v", *node)

		// fmt.Printf("%v", t.Nodes[p].Children)
		// put in end nodes map
		delete(t.EndNodes, p)
		t.EndNodes[n] = node
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

// find parent and move from one parent to the other
// first find greatest common ancestor
// actually this is probably like a dfs or bfs problem
