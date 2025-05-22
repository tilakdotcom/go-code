package main

import "fmt"

//for -> only construct in go for looping
func main() {
	//while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//inifinite loop
	// for {
	// 	fmt.Println("Hie man")
	// }

	//actual classic for loop

	for i := 0; i <= 10; i++ {

		if i == 2 {
			continue
		} else if i > 5 {
			break
		}

		fmt.Println(i)

	}

	//with range
	for i := range 5 {
		fmt.Println(i)
	}

}
