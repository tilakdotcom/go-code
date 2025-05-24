package main

import (
	"fmt"
	"time"
)

//channels are block

func deadLock() {
	messageChan := make(chan string)

	messageChan <- "ping" //blocking

	msg := <-messageChan

	fmt.Println(msg)
}

// sending channels's data
func printNum(nums chan int) {
	for num := range nums {
		fmt.Println("this is num ", num)
		time.Sleep(time.Second)
	}
}

// goroutine synchronizer
func task(done chan bool) {
	defer func() {
		done <- true
	}()

	fmt.Println("processing...")
}

// receiving channels's data
func sum(result chan int, nums ...int) {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	result <- total
}

func emailSending(emailChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()
	for email := range emailChan {
		fmt.Println("sending email to ", email)
		time.Sleep(time.Second)
	}
}

func selectCase() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	//assigning values
	go func() {
		chan1 <- 25
	}()
	go func() {
		chan2 <- "bhai"

	}()

	//printing
	for range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Println(chan1Val)
		case chan2Val := <-chan2:
			fmt.Println(chan2Val)

		}
	}
}

func main() {
	// deadLock()
	//sending channels data
	// nums := make(chan int)

	// go printNum(nums)

	// for range 2 {
	// 	nums <- rand.Intn(100)
	// }

	//receiving channels data

	// result1 := make(chan int)

	// go sum(result1, 32, 45, 134, 45, 23, 23, 43)

	// res := <-result1 //blocking

	// fmt.Println(res)

	//goroutine synchronizer
	done := make(chan bool)

	// go task(done)

	// <-done// for blocking the program

	//email sending
	emailChan := make(chan string, 100)

	for n := range 2 {
		emailChan <- fmt.Sprintf("%d@gmail.com", n)

	}

	go emailSending(emailChan, done)

	close(emailChan) //closing  channels

	<-done // for blocking the program

	//select Case
	selectCase()

}
