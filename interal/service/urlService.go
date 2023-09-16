package service

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type URLShortener struct {
	Origin string
	Short  string
	Ctime  int64
}

var dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=Asia/Taipei", os.Getenv("pgHost"), os.Getenv("pgPort"), os.Getenv("pgUser"), os.Getenv("pgPassword"), os.Getenv("pgDbname"))

type UrlService struct {
}

func NewUrlService() UrlService {
	return UrlService{}
}

func (service *UrlService) Generate(context *gin.Context) (string, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return "", err // to do error struct
	}
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	url := context.PostForm("url")

	var shortUrl string
	for i := 0; i < 3; i++ {
		shortUrl = service.getRandomStr(6)
		var count int64
		db.Table("`url-shortener`").Where("short = ?", shortUrl).Count(&count)
		if count == 0 {
			newURL := URLShortener{
				Short:  shortUrl,
				Origin: url,
				Ctime:  time.Now().Unix(),
			}
			if err := db.Table("url-shortener").Create(&newURL).Error; err == nil {
				return shortUrl, nil
			}
		}
	}

	return "", fmt.Errorf("Failed to generate URL too many times")
}

func (s *UrlService) GetQRcode(host string, shortUrl string) (string, error) {
	png, err := qrcode.Encode(host+"/"+shortUrl, qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}

	dataURI := "data:image/png;base64," + base64.StdEncoding.EncodeToString([]byte(png))
	return dataURI, nil
}

func (s *UrlService) getRandomStr(n int) string {
	a := make([]byte, n)
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range a {
		a[i] = letters[rand.Intn(len(letters))]
	}
	return string(a)
}

func (service *UrlService) GetShortUrl(context *gin.Context) ([]URLShortener, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	var urlShorteners []URLShortener
	db.Table("url-shortener").Find(&urlShorteners)
	return urlShorteners, nil
}

func (service *UrlService) GetOriginUrl(context *gin.Context) (string, error) {
	shortUrl := context.Param("shortUrl")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return "", err
	}
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	var result URLShortener

	if err := db.Table("url-shortener").Where("short = ?", shortUrl).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("Short URL not found")
		}
		return "", err
	}
	return result.Origin, nil

}
