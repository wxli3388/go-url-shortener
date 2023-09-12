package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/skip2/go-qrcode"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	router := gin.New()
	router.LoadHTMLGlob("./../web/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/list", list)
	router.GET("/:shortUrl", shortUrl)
	router.POST("/generate", generater)

	router.Run(":8080")
}

type URLShortener struct {
	Origin string
	Short  string
	Ctime  int64
}

var dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=Asia/Taipei", os.Getenv("pgHost"), os.Getenv("pgPort"), os.Getenv("pgUser"), os.Getenv("pgPassword"), os.Getenv("pgDbname"))

func list(context *gin.Context) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		context.HTML(http.StatusOK, "list.html", gin.H{
			"error": "Something went wrong...",
		})
	}
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	var urlShorteners []URLShortener
	db.Table("url-shortener").Find(&urlShorteners)
	// for i, v := range urlShorteners {
	// 	t := time.Unix(v.Ctime, 0)
	// 	urlShorteners[i].CtimeStr = t.Format("2006-01-02 15:04:05")
	// }
	context.HTML(http.StatusOK, "list.html", gin.H{
		"error": "",
		"data":  urlShorteners,
	})
}

func shortUrl(context *gin.Context) {
	shortUrl := context.Param("shortUrl")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"error": "Failed to connect to the database",
		})
		return
	}
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	var result URLShortener

	if err := db.Table("url-shortener").Where("short = ?", shortUrl).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Short url not found",
			})
			return
		}
	}
	context.Redirect(http.StatusMovedPermanently, result.Origin)
}

func generater(context *gin.Context) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"error": "Failed to connect to the database",
		})
		return
	}
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	url := context.PostForm("url")

	var shortUrl string
	for i := 0; i < 3; i++ {
		shortUrl = GetRandomStr(6)
		newURL := URLShortener{
			Short:  shortUrl,
			Origin: url,
			Ctime:  time.Now().Unix(),
		}

		if err := db.Table("url-shortener").Create(&newURL).Error; err != nil {
			// Handle the error
		} else {
			break
		}
	}

	if err != nil {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"error": err,
		})
		return
	}

	png, err := qrcode.Encode(context.Request.Host+"/"+shortUrl, qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}

	dataURI := "data:image/png;base64," + base64.StdEncoding.EncodeToString([]byte(png))

	context.HTML(http.StatusOK, "index.html", gin.H{
		"shortUrl": shortUrl,
		"qrcode":   template.URL(dataURI),
		"error":    "",
	})
}

func GetRandomStr(n int) string {
	a := make([]byte, n)
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range a {
		a[i] = letters[rand.Intn(len(letters))]
	}
	return string(a)
}
