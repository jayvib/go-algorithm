package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

func main() {
	q := queue.New()
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	for q.Len() != 0 {
		fmt.Print(q.Dequeue(), " ")
	}
}
