package main

import (
	"fmt"
	"shortUrl/db"
	"shortUrl/handler"
	"shortUrl/model"
)

func main() {
	err := db.Init()
	if err !=nil{
		return
	}
	var text string
	fmt.Printf("please Enter 1 or 2 : \n(1) to covert original link to short link \n(2) to covert short link to original link\n")
	fmt.Printf("menu is : ")
	_, err = fmt.Scanln(&text)
	if text == "1"{
		var input string
		fmt.Printf("please enter original link : ")
		_, err = fmt.Scanln(&input)
		shortUrl := handler.GenerateShortUrl()
		urlStore := &model.UrlStore{
			ShortUrl: shortUrl,
			FullUrl: input,
		}
		service := handler.Handler{}
		err = service.SaveShortUrl(urlStore)
		if err !=nil{
			fmt.Printf("error : %++v\n", err.Error())
		}else {
			fmt.Printf("short url is : %s",shortUrl)
		}
	}else if text == "2"{
		fmt.Printf("please enter short link : ")
		var input string
		_, err = fmt.Scanln(&input)
		service := handler.Handler{}
		result,err := service.GetOriginalUrl(input)
		if err != nil{
			fmt.Printf("error : %++v\n", err.Error())
		}else {
			fmt.Printf("original is : %s\n", result)
		}
	}else{
      fmt.Println("you enter wrong menu")
	}





}
