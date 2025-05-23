package main

import "fmt"

func addTwoNums(a int, b int) int {
	return a + b
}

func proccessAFunc(fn func(a int) int, b int) int {
	result := fn(3)
	return result * b
}

func proccessAFuncRetunFunc() func(a int) int {
	return func(a int) int {
		return a * 2
	}
}

//return more than one values
func returningMultipleValues() (string, string, string) {
	return "tilak", "developer", "golang"
}

func main() {
	result1 := addTwoNums(4, 2)
	fmt.Println(result1)

	//funtion inside func
	result2 := proccessAFunc(func(a int) int {
		return a + 2
	}, 3)

	fmt.Println(result2)

	//funtion returning funtions
	result3 := proccessAFuncRetunFunc()

	fmt.Println(result3(4))

	//getting more than one values
	name, work, _ := returningMultipleValues()

	fmt.Println(name, work)
}
