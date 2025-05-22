package main

import "fmt"

func main() {
	//arrays like c++
	//an empty uses there zeroed value by default
	// for example
	// int -> 0 on empty place
	// string -> "" on empty place
	//simple array
	arr := [4]int{}
	arr[3] = 4
	// fmt.Println(arr)
	fmt.Println(len(arr))

	//declare array in single line
	arr2 := [5]int{1, 4, 5, 6}
	fmt.Println(arr2)

	// 2D arrays
	arr3 := [3][3]int{{1, 2, 4}, {4, 2, 5}}
	fmt.Println(arr3)

	// -fixed size, that is predefined and predictable
	// -memory optimization
	// -constant time access
}
