package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"html/template"
		"encoding/xml"
		"sync")

var wg sync.WaitGroup

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

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

type NewsAggPage struct{
	Title string
	//News string
	News map[string]NewsMap
}

func newsRoutine(channel chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ :=  http.Get(Location) 
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	channel <- n
}

func newsAggHandler(writer http.ResponseWriter, request *http.Request){

	//resp, _ :=  http.Get("https://www.thehindu.com/") // _ is for error
	resp, _ :=  http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") 
	bytes, _ := ioutil.ReadAll(resp.Body)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	resp.Body.Close()

	var s SitemapIndex
	
	news_map := make(map[string]NewsMap)
	xml.Unmarshal(bytes, &s)

	queue := make(chan News, 30)
	//fmt.Println(s.Locations)
	for _, Location := range s.Locations { // _ is index value
		//fmt.Printf("%s\n", Location)
		wg.Add(1)
		go newsRoutine(queue, Location)
	}

	wg.Wait()
	close(queue)

	for elem := range queue {
		for idx, _ := range elem.Keywords {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}	
	}
	

	//p := NewsAggPage{Title: "Amazing news aggregator", News:"some news"}
	p := NewsAggPage{Title: "Amazing news aggregator", News: news_map}
	t, _ := template.ParseFiles("basictemplating.html")
	fmt.Println(t.Execute(writer, p))
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}