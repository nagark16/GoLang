package main

import ("fmt"
		"net/http")

func index_handler(writer http.ResponseWriter, request *http.Request){ //* mean reading through http request
	//fmt.Fprintf(writer, "Go is neat")
	fmt.Fprintf(writer, "<h1>Go is neat</h1>")
	fmt.Fprintf(writer, "<p>Go is neat</p>")
	fmt.Fprintf(writer, "<p>%s is %s</p>", "Go", "<strong>neat</strong>")

	fmt.Fprintf(writer,`<h1>Go is neat</h1>
				<p>Go is neat</p>
		`)
}

func about_handler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Naga here")
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}