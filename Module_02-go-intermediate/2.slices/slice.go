package main

import (
	"fmt"
	"slices"
)

// slice is like Vector in c++
//it is next version of array with lots og methods like map filter
//most used construct in go

func main() {
	// uninitial slice is nil
	// nil is same as NULL
	//normal  slice -> Array
	var slice1 []int

	fmt.Println(slice1 == nil)
	fmt.Println(len(slice1))

	//slice with default values
	var slice2 = make([]int, 3)
	fmt.Println(slice2 == nil)

	fmt.Println(slice2)

	//capacity
	var slice3 = make([]int, 3, 5)
	// make.1 is for defining type of slice
	// make.2 is for defining initial  size
	// make.3 is for defining initial capacity

	//capacity -> maximum numbers of  elements can fit
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 3)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 4)
	slice3 = append(slice3, 3)
	fmt.Println(len(slice3))
	fmt.Println(slice3)
	fmt.Println(cap(slice3))

	//ideal way to use make is using 0 on the size prop
	var slice4 = make([]int, 0, 5)
	slice4 = append(slice4, 1)
	slice4 = append(slice4, 2)
	fmt.Println(slice4)

	//shorthand syntax
	slice5 := []int{} //empty slice
	fmt.Println(slice5)
	slice5 = append(slice5, 1)
	slice5 = append(slice5, 2)
	slice5 = append(slice5, 4)
	// slice1[0] = 5 //throw error coz it coesn't have inital size
	slice3[0] = 34 // but is will work

	fmt.Println(slice5)

	//copy funtions
	var slice6 = make([]int, 0, 5)

	slice6 = append(slice6, 1)
	slice6 = append(slice6, 2)
	slice6 = append(slice6, 3)

	var copiesSlice = make([]int, len(slice6))

	//coping array
	copy(copiesSlice, slice6)
	fmt.Println(slice6, copiesSlice)

	//slice oparator
	slice7 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice7[1:3])
	fmt.Println(slice7[:3])
	fmt.Println(slice7[:])
	fmt.Println(slice7[2:])

	//slices package
	slice8 := []int{1, 2, 3, 4, 5, 6}
	checkingSlice := []int{1, 2, 3, 4, 6}

	fmt.Println(slices.Equal(slice8, checkingSlice))

	//2D slice
	var slice2D = [][]int{{1, 3, 4}, {1, 2, 3, 4, 5}}

	fmt.Println(slice2D)

}
