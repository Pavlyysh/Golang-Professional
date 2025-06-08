package main

import "fmt"

// Можно взять пакет container/list
type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)
	addNode(root, 100)
	traverse(root)
	fmt.Println(deleteNode(root, 1))

	traverse(root)
}

func deleteNode(t *Node, v int) bool {
	if t == nil {
		fmt.Println("Empty list")
		return false
	}

	if t.Value == v {
		root = t.Next
		return true
	}

	current := t
	for current.Next != nil {
		if current.Next.Value == v {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}

	return false
}

func addNode(t *Node, v int) int {
	if root == nil {
		t := &Node{v, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("already exists", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
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

func lookupNode(t *Node, v int) bool {
	if root == nil {
		t := &Node{v, nil}
		root = t
		return false
	}

	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list")
		return 0
	}

	sum := 0
	for t != nil {
		sum++
		t = t.Next
	}

	return sum
}
