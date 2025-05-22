package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 4

	switch i {
	case 1:
		fmt.Println("case 1")

	case 2:
		fmt.Println("case 2")
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("default case")
	}

	// multiple switch condition
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's Weekend")
	default:
		fmt.Println("it's working day ")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("It is INT")
		case string:
			fmt.Println("It is  STRING ")
		case bool:
			fmt.Println("It is  BOOL")
		default:
			fmt.Println("It is unknoun", i)

		}

	}

	whoAmI("tilak")
	whoAmI(4)
	whoAmI(true)
	whoAmI(4.234)
}
