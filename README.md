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