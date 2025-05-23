package main

import "fmt"

type Slice[T any] struct {
	element []T
}

//normal
// func printSlice[T any](items []T) {
// 	for _,item := range items {
// 		fmt.Println(item)
// 	}
// }
//will work the same
func printSliceWithComparable[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printSliceWithLimitation[T string | int](items []T) {
// 	for _,item := range items {
// 		fmt.Println(item)
// 	}
// }
//similar
// func printSliceWithInterface[T interface{}](items []T) {
// 	for _,item := range items {
// 		fmt.Println(item)
// 	}
// }

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	printSliceWithComparable(nums)

	arr1 := Slice[string]{
		element: []string{"Go", "java", "javascript", "c", "nodejs", "c++"},
	}
	printSliceWithComparable(arr1.element)

}
