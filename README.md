# GoLang
GoLang practice

Quick reference to my GoLang practice

1. For installation https://www.admfactory.com/how-to-install-golang-1-10-on-ubuntu/

2. To run go program 
```
	go run basics.go
```
3. To fix go file style mistakes
```
	gofmt basics.go
```
4. using *make* we can create following things
	* channel
	* map with space preallocated
	* slice with space preallocated or with len != cap 
5. Arrays
```
	var buffer [256]byte
	var slice []byte = buffer[100:150] or var slice = buffer [100:150] or slice := buffer[100:150] //100 inclusive. 150 exclusive
```
internally slice will create
```
	type sliceHeader struct {
	    Length        int
	    Capacity      int
	    ZerothElement *byte
	}

	slice := sliceHeader{
	    Length:        50,
	    ZerothElement: &buffer[100],
	}
```
6. The slice created by
```
array[0:0]
```

has length zero (and maybe even capacity zero) but its pointer is not nil, so it is not a nil slice. An empty slice can grow (assuming it has non-zero capacity), but a nil slice has no array to put values in and can never grow to hold even one element. That said, a nil slice is functionally equivalent to a zero-length slice, even though it points to nothing. It has length zero and can be appended to, with allocation

7. fmt.Printf
	%x - hexa decimal
	%q - quoted
	%+q - The + flag causes the output to escape not only non-printable sequences, but also any non-ASCII bytes, all while interpreting UTF-8.

8. calling a method with go will create a thread.
9. to have communicatin between threads we can use channels, which we can create via 
```
	intChannel := make(chan int) // creates a channel which transports integer data type
	intChannel := make(chan int, 2) // creates a channel with capacity of 2. Also, Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty
	intChannel, ok := make(chan int, 2) // ok is false if there are no more values to receive and the channel is closed by sender
```
10. Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop. 
11. Accessing struct pointer is different
```
	type Vertex struct {
		X int
		Y int
	}

	func main() {
		v := Vertex{1, 2}
		p := &v
		fmt.Println(p) // prints &{1 2}
		p.X = 1e9
		fmt.Println(p) // prints &{1000000000 2}
		fmt.Println(v.X) // prints 1000000000
		fmt.Println(p.X) // prints 1000000000 // we can access field of struct pointer p filed via p.X instead of (*p).X
		fmt.Println((*p).X) // prints 1000000000
		//fmt.Println(*p.X) // error
	}

```
12. Struct literals: We can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.) The special prefix & returns a pointer to the struct value. 
```
	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	func main() {
		fmt.Println(v1, p, v2, v3) // {1 2} &{1 2} {1 0} {0 0}
	}

```
13. Function Closures: A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
```
	func adder() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	func main() {
		pos, neg := adder(), adder()
		for i := 0; i < 10; i++ {
			fmt.Println(
				pos(i),
				neg(-2*i),
			)
		}
	}
```