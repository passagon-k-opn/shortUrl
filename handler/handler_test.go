package handler

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"shortUrl/model"
	"testing"
)

func getSession() (*gorm.DB,sqlmock.Sqlmock)  {
	db,mock,err:= sqlmock.New()
    convey.So(err,convey.ShouldBeNil)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	convey.So(err,convey.ShouldBeNil)
	return gormDB,mock

}
func TestGenerateShortUrlAndSave(t *testing.T) {
	convey.Convey("test..",t, func() {
		db,mock := getSession()
		service := Handler{DB: db}
		convey.Convey("GenerateShortUrlAndSave Success", func() {
			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO \"url\".\"url_store\"").WillReturnResult(sqlmock.NewResult(1,1))
			mock.ExpectCommit()
			shortLink := GenerateShortUrl()
			convey.So(shortLink,convey.ShouldNotBeEmpty)
			urlStore:=&model.UrlStore{}
			err := service.SaveShortUrl(urlStore)
			convey.So(err,convey.ShouldBeNil)
		})
		convey.Convey("GenerateShortUrlAndSave Fail", func() {
			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO \"url\".\"url_store\"").WillReturnError(errors.New("cannot database"))
			mock.ExpectCommit()
			shortLink := GenerateShortUrl()
			convey.So(shortLink,convey.ShouldNotBeEmpty)
			urlStore:=&model.UrlStore{}
			err := service.SaveShortUrl(urlStore)
			convey.So(err,convey.ShouldNotBeNil)
		})
	})
}
func TestQueryOriginalUrl(t *testing.T) {
	convey.Convey("test..",t, func() {
		db,mock := getSession()
		service := Handler{DB: db}
		convey.Convey("QueryOriginalUrl Success", func() {
			//mock.ExpectBegin()
			mock.ExpectQuery("^SELECT (.*) FROM \"url\".\"url_store\"").WithArgs("test").WillReturnRows(sqlmock.NewRows(
				[]string{"short_url","full_url"}).AddRow("test","originalUrl"))
			//mock.ExpectCommit()
			result,err:= service.GetOriginalUrl("test")
			convey.So(err,convey.ShouldBeNil)
			convey.So(result,convey.ShouldEqual,"originalUrl")
		})
		convey.Convey("QueryOriginalUrl Fail", func() {
			mock.ExpectQuery("^SELECT (.*) FROM \"url\".\"url_store\"").WithArgs("test").WillReturnError(errors.New("cannot database"))
			result,err:= service.GetOriginalUrl("test")
			convey.So(err,convey.ShouldNotBeNil)
			convey.So(result,convey.ShouldEqual,"")
		})
	})
}
