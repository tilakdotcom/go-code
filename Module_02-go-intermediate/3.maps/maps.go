package main

import (
	"fmt"
	"maps"
)

//maps -> maps used to store data values in key: pairs

func main() {
	//ceating maps
	var map1 map[string]string = map[string]string{
		"name": "Tilak",
		"work": "development",
		"rn":   "learning Go",
	}

	// fmt.Println(len(map1))
	// fmt.Println(map1)
	//get by key
	fmt.Println(map1["name"], map1["work"])
	fmt.Println(map1["ash"])
	//imp: if key is not exists in the map then it returns  zeroed value

	// some oparations
	map2 := make(map[string]int)
	map2["age"] = 20
	map2["rank"] = 2
	map2["exp"] = 3

	//delete a key with its value
	delete(map2, "exp")

	//clear whole maps
	// clear(map2)

	fmt.Println(len(map2))
	fmt.Println(map2)

	// to check is item exists or not
	map3 := map[string]string{
		"name": "anshu",
		"work": "development",
		"rn":   "learning Go",
	}

	// v, ok := map3["name"] //exits
	v, ok := map3["ash"] //not exists
	//v state for keys values
	fmt.Println(v)
	if ok {
		fmt.Println("data is exists")
	} else {
		fmt.Println("data is not exists")
	}

	// maps
	map4 := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}

	copiedMap := map[int]string{
		1: "firstt",
		2: "second",
		3: "third",
	}

	//using maps
	fmt.Println(maps.Equal(map4, copiedMap))

}
