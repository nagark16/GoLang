package main

import("time"
		"fmt"
		"sync")

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
	//wg.Done()
}

func say(s string){
	//defer wg.Done()
	defer cleanup()
	for i:=0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
		/*if i == 2 {
			panic("gone wrong")
		}*/
	}
	//wg.Done() //without this the go routines will be in deadlock
}

func foo(){
	defer fmt.Println("Done")
	defer fmt.Println("Done2")
	fmt.Println("whats going on?")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	/*go say("hey") // this will run as a thread
	go say("there")
	time.Sleep(time.Second)// without this sleep the main may close before above go routines run*/

	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")
	wg.Wait()

	wg.Add(1)
	go say("hi")
	wg.Wait()

	foo()
}