package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
[5 5]int == array
[]int == slice(dynamic array)

*/

type SitemapIndex struct {
	Location []Location `xml:"sitemap"` //Make sure Location is capitalized becuase it needs to be exported

}

type Location struct {
	Loc string `xml:"loc"`
}

func main() {
	// http.HandleFunc("/", index_handler)
	// http.ListenAndServe(":8000", nil)

	resp, _ := http.Get("https://www.washingtonpost.com/business/2021/09/29/schools-supply-chain-crisis/?itid=hp-top-table-main")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	resp.Body.Close()

	fmt.Printf(string_body)

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

}

// func index_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")

// 	fmt.Fprintf(w, `<p>Go is fast!</p>
// 		<p>Hello sir</p>
// 	`)
// }

// Checkout out http://localhost:8000/ after running
