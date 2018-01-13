package main

import (
	"fmt"

	"stacker/stack"
)

func main() {
	var stack stack.Stack

	stack.Push("hello")
	stack.Push(54)
	stack.Push([]string{"I", "am", "Salar"})
	stack.Push(-652)

	for {
		item, err := stack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
