package model

type UrlStore struct {
	ShortUrl string `gorm:"primaryKey;column:short_url"`
	FullUrl string   `gorm:"column:full_url"`
}
