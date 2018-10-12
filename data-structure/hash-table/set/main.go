package main

import (
	"fmt"
	"github.com/golang-collections/collections/set"
)

func main() {
	st := set.New()
	st.Insert(1)
	st.Insert(2)

	fmt.Println(st.Has(1))
	fmt.Println(st.Has(3))
	fmt.Println(st.Len())

	st.Do(func(data interface{}){
		if i, ok := data.(int); ok {
			fmt.Println(i + 1)
		}
	})

	fmt.Println(st)
}

