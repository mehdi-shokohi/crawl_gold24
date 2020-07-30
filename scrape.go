package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/valyala/fasthttp"
	"log"
	"regexp"
	"strings"
)

func getGeram24(newPrice chan <- string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Error Occurred", r)
		}
	}()
	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("tbody[class='table-padding-lg']", func(e *colly.HTMLElement) {
		 e.ForEach("tr", func(i int, element *colly.HTMLElement)  {
			ok,_:= regexp.MatchString(" فعلی",element.Text)
		 	if ok{
				//fmt.Println(element.Text)
				price :=  strings.Replace(element.DOM.Find("td[class=text-left]").Text(),",","",-1)
				newPrice <-price

			}

		})

	})
	er:=c.Visit("https://www.tgju.org/profile/geram24")
	if er!=nil{
		panic(er)
	}
}


func main(){

	priceChan:=make(chan string)

	FHandler:=new(WebServerHandler)
	FHandler.pchan=priceChan
	er:=fasthttp.ListenAndServe(":8181", FHandler.RequestHandler)
	if er!=nil{
		log.Fatal("Error in Starting Server ")
	}
}