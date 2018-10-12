package main

import "fmt"

type Set map[interface{}]bool

func (s *Set) Add(key interface{}) {
	(*s)[key] = true
}

func (s *Set) Find(key interface{}) bool {
	return (*s)[key]
}

func (s *Set) Remove(key interface{}) {
	delete(*s, key)
}

func (s *Set) Size() int {
	return len(*s)
}

func main() {
	set := make(Set)

	set.Add("a")
	set.Add("b")
	set.Add(1)
	fmt.Println(set.Size())
	fmt.Println("is 'a' exist?", set.Find("a"))
	fmt.Println("is 'c' exist?", set.Find("c"))
	fmt.Println("is 'b' exist?", set.Find("b"))
	fmt.Println("is 1 exist?", set.Find(1))
}
