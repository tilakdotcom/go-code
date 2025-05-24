package main

import "fmt"

//enumerated Type

type StatusType string

const (
	Shipped StatusType = "shipped"
	Confirm StatusType = "confirm"
	Prepaid StatusType = "prepaid"
	Cancel  StatusType = "cancel"
)

func changeStatus(status StatusType) {
	fmt.Println("changin status to ", status)
}

func main() {
changeStatus(Cancel)
changeStatus(Confirm)
changeStatus(Prepaid)
changeStatus(Shipped)
changeStatus(Cancel)
}
