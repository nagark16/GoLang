package main

import ("fmt"
		"math"
		"math/rand")

/*
	2 types of methods	
		1. value receivers
		2. pointer receivers
*/

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct{
	gas_pedal uint16 // min 0 max 65535
	brake_pedal uint16
	steering_wheel int16 //-32k to +32k
	top_speed_kmh float64
}

//value receiver methods for car struct
func (c car) kmh() float64{
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

func (c car) mph() float64{
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

//pointer receiver
func (c *car) new_top_speed (newspeed float64){
	c.top_speed_kmh = newspeed
}

func newer_top_speed(c car, speed float64) car{
	c.top_speed_kmh = speed
	return c
}

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


	a_car := car{gas_pedal: 22341, 
				brake_pedal: 0, 
				steering_wheel: 12561, 
				top_speed_kmh: 225.0}

	b_car := car{22341, 0, 12561, 225.0}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(b_car.top_speed_kmh)

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car = newer_top_speed(a_car, 500)
}