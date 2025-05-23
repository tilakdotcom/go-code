package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int) {
	fmt.Println("doing task ", id)
}

func goroutineWithSleep() {
	for i := range 10 {
		go task(i + 1)

		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}

func goroutineWithWG(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	task(i)
}

func main() {
	// goroutineWithSleep()
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go goroutineWithWG(i, &wg)
	}
	wg.Wait()

}
