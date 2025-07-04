package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {
	tree := create(10)
	fmt.Println("the root value is:", tree.Value)
	traverse(tree)
	fmt.Println()
	insert(tree, -10)
	insert(tree, 11)
	traverse(tree)
	fmt.Println()
	fmt.Println("the root value is:", tree.Value)
}

func traverse(t *Tree) {
	if t == nil {
		return
	}

	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < n*2; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}
