package db

import (
	"fmt"
	"gorm.io/gorm"
	"shortUrl/model"
)

func InsertUrlStore(engine *gorm.DB,urlStore *model.UrlStore) error  {
	if err:= engine.Table("url.url_store").Create(&urlStore).Error; err !=nil{
		fmt.Printf("InsertUrlStore error : %++v\n", err)
		return err
	}
	return nil
}
func QueryUrlStore(engine *gorm.DB,shortUrl string) (url string,err error)  {
	urlStore := &model.UrlStore{}
	if err:= engine.Table("url.url_store").Where("short_url = ?",shortUrl).First(urlStore).Error; err !=nil{
		fmt.Printf("QueryUrlStore error : %++v\n", err)
		return "",err
	}
	return urlStore.FullUrl,nil
}