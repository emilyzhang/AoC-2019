package tree

import (
	"fmt"
)

// Node .
type Node struct {
	Name     string
	Parent   *Node
	Children []*Node
}

// Tree .
type Tree struct {
	Nodes    map[string]*Node
	EndNodes map[string]*Node
}

// NewTree .∂
func NewTree() *Tree {
	t := Tree{
		Nodes:    make(map[string]*Node),
		EndNodes: make(map[string]*Node),
	}
	return &t
}

// CountOrbits .
func (t *Tree) CountOrbits() int {
	var count int
	for _, node := range t.Nodes {
		n := node
		for n.Parent != nil {
			count++
			n = n.Parent
		}
	}
	return count
}

// PrettyPrint .
func (t *Tree) PrettyPrint() string {
	s := ""
	for _, node := range t.Nodes {
		if node.Parent != nil {
			s += "\n" + node.Parent.Name + ")" + node.Name + " children: "
		} else {
			s += "\n" + "NO PARENT" + ")" + node.Name + " children: "
		}
		for _, child := range node.Children {
			if child != nil {
				s += child.Name + ","
			}
		}
	}
	return s
}

// Distance .
func (t *Tree) Distance(start, end string) int {
	seen := make(map[string]bool)
	from := t.Nodes[start].Parent
	to := t.Nodes[end].Parent
	println(from.Name, to.Name)
	minDistance := 100000000000000000
	type pair struct {
		name     string
		distance int
	}
	q := make([]pair, 0)
	q = append(q, pair{from.Name, 0})

	var p pair
	var curr string
	var currDistance int
	var currNode *Node
	for len(q) > 0 {
		p, q = q[0], q[1:]
		fmt.Printf("%v", p)
		curr = p.name
		currDistance = p.distance
		if curr == to.Name {
			if currDistance < minDistance {
				print(currDistance)
				minDistance = currDistance
			}
		} else {
			// add parent and all children, unless already seen
			currNode = t.Nodes[curr]
			if currNode.Parent != nil {
				if _, ok := seen[currNode.Parent.Name]; !ok {
					q = append(q, pair{currNode.Parent.Name, currDistance + 1})
					seen[currNode.Parent.Name] = true
				}
			}
			for _, child := range currNode.Children {
				if _, ok := seen[child.Name]; !ok {
					q = append(q, pair{child.Name, currDistance + 1})
					seen[child.Name] = true
				}
			}
		}
	}
	return minDistance
}
