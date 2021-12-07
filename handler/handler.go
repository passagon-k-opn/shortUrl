package handler

import (
	"gorm.io/gorm"
	"math/rand"
	"shortUrl/constant"
	"shortUrl/db"
	"shortUrl/model"
	"time"
)

type Handler struct {
	DB *gorm.DB
}

func (t *Handler) initDB() {
	if nil ==t.DB{
		t.DB = db.EngineDB
	}
}
func GenerateShortUrl() string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, constant.LengthOfUrl)
	for i := range b {
		b[i] = constant.CHARSET[seededRand.Intn(len(constant.CHARSET))]
	}
	shortUrl := constant.HostnameShortUrl + string(b)
	return shortUrl
}
func (t *Handler) SaveShortUrl(urlStore *model.UrlStore) error {
	t.initDB()
	engineDb := t.DB
	err := db.InsertUrlStore(engineDb, urlStore)
	if err != nil {
		return err
	}
	return nil
}
func (t *Handler)GetOriginalUrl(shortUrl string) (string, error) {
	t.initDB()
	engineDb := t.DB
	originalUrl, err := db.QueryUrlStore(engineDb, shortUrl)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}
