package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count = count + 1
		return count
	}
}

func main() {
	//closure example
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
