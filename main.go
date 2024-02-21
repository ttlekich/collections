package main

import "fmt"

func main() {
	l := NewSLList[string]()

	l.Push("Hello")
	l.Push("World")

	out, err := l.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*out)
}
