package main

import "fmt"

func main() {
	// if Else
	age := 16
	if age >= 18 {
		fmt.Println("user is Adult")
	} else if age >= 12 {
		fmt.Println("user is Teenager")
	} else {
		fmt.Println("user is Kid")
	}

	// OR and AND
	role := "admin"
	hasAccess := true

	if role == "admin" && hasAccess {
		fmt.Println("user is Admin and has access")
	} else if role == "admin" && !hasAccess {
		fmt.Println("user is admin but not has access")
	} else {
		fmt.Println("user not has access")
	}

	if role == "admin" || hasAccess {
		fmt.Println("either user is admin or has Access may be both")
	}
	//can create var in if else block and access
	if age := 15; age >= 18 {
		fmt.Println("user is an Adult", age)
	} else if age >= 12 {
		fmt.Println("user is Teenager", age)
	}

	fmt.Println(age)

	// go does not have ternary, you will have to use normal if else
}
