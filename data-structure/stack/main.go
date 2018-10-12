package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {
	s := stack.New()
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push("six")
	s.Push("seven")

	for s.Len() != 0 {
		fmt.Println(s.Pop())
	}
}
