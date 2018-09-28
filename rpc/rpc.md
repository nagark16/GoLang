1. To make a method to be remotly accessible the following requirements should be met
	a. The method’s type is exported (ex: type Export int)
	b. The method is exported (by capitalising method name)
	c. The method has two arguments, both exported (or builtin types).
	d. The method’s second argument is a pointer (by declaring return type)
	e. The method has return type error (after method signature closing brackets)
///
	type Task int (a)
	func (t *Task)(a) MakeToDo(b)(todo ToDo, reply *ToDo)(d) error(e) {
		todoSlice = append(todoSlice, todo)
		*reply = todo
		return nil
	}
///
2. To call the above method
///
	task := new(Task)
	var err error
	finishApp := ToDo{"Finish App", "Started"}
	var makeReply ToDo
	err = task.MakeToDo(finishApp, &makeReply)
	if err != nil {
		log.Fatal("Issue making ToDo: ", err)
    }
///
3. Exporting method should have 2 arguments. From documentation
```
	The method’s first argument represents the arguments provided by the caller; the second argument represents the result parameters to be returned to the caller. The method’s return value, if non-nil, is passed back as a string that the client sees as if created by errors.New. If an error is returned, the reply parameter will not be sent back to the client.
```
4. Server which exposes this methods acts as RPC server
///
	task := new(Task)
	
	// Publish the receivers methods
	err := rpc.Register(task)
	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
	}
	
	// Register a HTTP handler
	rpc.HandleHTTP()
	
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", 1234)
	
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
///
5. Client to make calls to server
///
	var err error
	var reply ToDo
	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	finishApp := ToDo{"Finish App", "Started"}
	makeDinner := ToDo{"Make Dinner", "Not Started"}
	walkDog := ToDo{"Walk the dog", "Not Started"}

	client.Call("Task.MakeToDo", finishApp, &reply)
	client.Call("Task.MakeToDo", makeDinner, &reply)
///