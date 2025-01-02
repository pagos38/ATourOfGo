package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

func Walk(t *Tree, ch chan int) {
	defer close(ch)
	walk(t, ch)
}

func walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func Same(t1, t2 *Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		switch {
		case !ok1, !ok2:
			return ok1 == ok2
		case v1 != v2:
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(1), New(2)))
}
