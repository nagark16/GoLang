package main

import ("fmt"
		"math"
		"math/rand")

//func add(x float64, y float64) float64{
func add(x , y float64) float64{
	return x + y
}

func foo(){
	fmt.Println("welcome to go")
	fmt.Println("square root of 4 is ", math.Sqrt(4))
	fmt.Println("random number from 1 - 100 ", rand.Intn(100))
}

const x int = 5

func multiple(a,b string) (string,string){
	return a,b
}

func main() {
	foo()

	/*var num1 float64 = 5.6
	var num2 float64 = 9.5*/

	//var num1,num2 float64 = 5.6, 9.5 // var is required if we are declaring outsie a function

	num1,num2  := 5.6, 9.5 //here go will figure out the data type and that too float64

	fmt.Println(add(num1,num2))

	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1,w2))

	var a int = 62
	var b float64 = float64(a)

	x := a // type inference. x will be int
	fmt.Println(b)
	fmt.Println(x)
	y := 15
	c := &y // memory address
	fmt.Println(c)
	fmt.Println(*c)
	*c = 5
	fmt.Println(y)
	*c = *c**c
	fmt.Println(y)
}