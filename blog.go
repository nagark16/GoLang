package main

import("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")

//[5]type == array
//[]type == slice

/*type SitemapIndex struct{
	Locations []Location `xml:"sitemap"`
}

type Location struct{
	Loc string `xml:"loc"`
}

func (l Location) String() string{
	return fmt.Sprintf(l.Loc)
}*/

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

func main() {
	//resp, _ :=  http.Get("https://www.thehindu.com/") // _ is for error
	resp, _ :=  http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") 
	bytes, _ := ioutil.ReadAll(resp.Body)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	resp.Body.Close()

	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)
	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)
	for _, Location := range s.Locations { // _ is index value
		//fmt.Printf("%s\n", Location)
		resp, _ :=  http.Get(Location) 
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for idx, data := range news_map {
		fmt.Println("\n\n\n",idx)
		fmt.Println("\n",data.Keyword)
		fmt.Println("\n",data.Location)
	}
}