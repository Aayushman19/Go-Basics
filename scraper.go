package main

import (
	"fmt"
	"os"
	"github.com/gocolly/colly"
)

func main(){
	args := os.Args
	url := args[1]
	collector := colly.NewCollector()
	
}