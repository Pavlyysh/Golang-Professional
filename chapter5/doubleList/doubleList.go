package main

import "fmt"

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

var root = new(Node)

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, 0)
	traverse(root)
	addNode(root, 100)
	fmt.Println("Size:", size(root))
	traverse(root)
	reverse(root)
}

func addNode(t *Node, v int) int {
	if root == nil {
		t := &Node{v, nil, nil}
		root = t
		return 0
	}
	if t.Value == v {
		fmt.Println("already exists")
		return -1
	}
	if t.Next == nil {
		temp := t
		t.Next = &Node{v, temp, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Prev != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Prev
	}
	fmt.Printf("%d ->", temp.Value)
	fmt.Println()
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("Empty list")
		return 0
	}

	sum := 0
	for t != nil {
		sum++
		t = t.Next
	}
	return sum
}

func lookupNode(t *Node, v int) bool {
	if t == nil {
		return false
	}

	if t.Value == v {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}
