package main

import ("fmt"
		"sync"
		)


var wg sync.WaitGroup

func foo(channel chan int, somevalue int) {
	defer wg.Done()
	channel <- somevalue * 5 // here <- mean to send a value on a channel
}

func main() {
	fooVal := make(chan int, 10) // 10 mean 10 items to buffer
	/*go foo(fooVal, 5)
	go foo(fooVal, 3)*/

	/*v1 := <- fooVal  // here <- mean to receive a value
	v2 := <- fooVal */

	/*v1, v2 := <- fooVal, <- fooVal
	fmt.Println(v1, v2)*/

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i) //without 'go' here it will be deadlock
	}

	wg.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}

}