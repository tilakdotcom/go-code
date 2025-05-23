package main

import "fmt"

func sum(nums ...int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum = totalSum + num
	}
	return totalSum
}

func console(props ...interface{}) {
	fmt.Println(props...)
}

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7}

	result := sum(slice1...)
	console(result)

}
