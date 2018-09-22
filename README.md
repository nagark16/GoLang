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