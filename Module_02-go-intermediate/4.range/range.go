package main

import "fmt"

// more about range

func main() {
	slice1 := []string{"men", "women", "hello", "me"}
	//bacis impmenentation
	// for i := 0; i < len(slice1); i++ {
	// 	fmt.Println(slice1[i])
	// }

	//using range
	for i, v := range slice1 {
		fmt.Println(i, v) // i i stand for index
	}

	map1 := map[int]string{
		1: "firstt",
		2: "second",
		3: "third",
	}

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//range on string
	//unicode coe point rune ASCII
	//i is starting byte of runa
	//255 -> 1bytes
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
