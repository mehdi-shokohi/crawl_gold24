package main

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)




type PriceData struct {
	Result bool `json:"result"`

		Data struct {
			Prices struct {
				Geram24 struct {
					Current string `json:"current"`
				} `json:"geram24"`
			} `json:"prices"`
		} `json:"data"`
}
type WebServerHandler struct{
	pchan chan  string
}
func (self * WebServerHandler) RequestHandler(ctx *fasthttp.RequestCtx){
	go getGeram24(self.pchan)

	select {
	case <-time.After(4000*time.Millisecond):
		responseData:=new(PriceData)
		responseData.Data.Prices.Geram24.Current="-1"
		responseData.Result=false
		resFinal,_:=json.Marshal(responseData)
		ctx.Response.Header.Set("Content-type","application/json")
		_,er:= ctx.Write(resFinal)
		if er!=nil{
			log.Fatal("Error In Send message To Client")
		}
	case price:= <- self.pchan :
		responseData:=new(PriceData)
		//d,_:=strconv.Atoi(price)
		responseData.Data.Prices.Geram24.Current=price
		responseData.Result=true
		resFinal,_:=json.Marshal(responseData)

		ctx.Response.Header.Set("Content-type","application/json")
		_,er:= ctx.Write(resFinal)
		if er!=nil{
			log.Fatal("Error In Send message To Client")
		}
		//fmt.Println("Price : ",price)

	}

}