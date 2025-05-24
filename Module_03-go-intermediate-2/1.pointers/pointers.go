package main

import "fmt"

//by Value only, It has only Distinct Copy of actual var
func copyActualVar(num int) {
	num = 5 //won't effect the actual variable
	fmt.Println("in copyActualVar props is ", num)
}

//by reference, it will have  the actual value n memory address of variable
func referenceOfVariable(num *int) {
	*num = 27

	fmt.Println("variable in referenceOfVariable ", *num)

}

func main() {
	num := 1

	copyActualVar(num)
	fmt.Println("actual var ", num)

	fmt.Println("memory Address ", &num)

	referenceOfVariable(&num)

	fmt.Println("actual var ", num)

}
