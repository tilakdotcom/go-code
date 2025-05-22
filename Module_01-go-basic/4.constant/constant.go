package main

import "fmt"

// work := "development" // can not use this method here
const age int = 20

var lang string = " GO"

func main() {
	const name string = "Tilak Singh" //constant

	//can not asign new valaue to name coz it is using const
	// name = "anshu"
	fmt.Println(name)

	// constant grouping 
	const (
		port = 5000
		host = "localHost"
	)

	fmt.Println(port, host)


}
