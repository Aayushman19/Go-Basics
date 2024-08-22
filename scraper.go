package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main(){
	args := os.Args
	url := args[1]
	collector := colly.NewCollector()

	dict, err := os.Create("testfile.txt")

	if err != nil{
		log.Panic(err)
	}

	line, err := dict.WriteString("File starts here")

	if err != nil{
		panic(err)
	}

	fmt.Println(line)

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error occured ", err)
	})

	var element string

	fmt.Println("Enter the html tag you want to scrape")
	fmt.Scan(&element)


	collector.OnHTML(element, func(h *colly.HTMLElement) {
		var str = h.DOM.Text()
		temp, err := dict.WriteString(str)
		if err != nil{
			panic(err)
		}
		fmt.Println(temp)
	})

	collector.Visit(url)
}