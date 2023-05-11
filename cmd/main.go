package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.New()
	router.LoadHTMLGlob("./../web/*")
	router.GET("/", func(c *gin.Context) {
		fmt.Println()
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/:shortUrl", ShortUrl)
	router.POST("/generate", Generater)

	router.Run(":3388")
}

type Url struct {
	Short  string
	Origin string
}

func ShortUrl(context *gin.Context) {
	shortUrl := context.Param("shortUrl")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", os.Getenv("pgHost"), os.Getenv("pgPort"), os.Getenv("pgUser"), os.Getenv("pgPassword"), os.Getenv("pgDbname"))
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(`SELECT origin FROM "url-shortener" WHERE short=$1`)
	if err != nil {
		// log.Fatal(err)
	}
	defer stmt.Close()

	var url Url
	err = stmt.QueryRow(shortUrl).Scan(&url.Origin)
	if err != nil {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"error": "Doesn't find this short url",
		})
		return
	}

	fmt.Println(url)
	context.Redirect(http.StatusMovedPermanently, url.Origin)
}

func Generater(context *gin.Context) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", os.Getenv("pgHost"), os.Getenv("pgPort"), os.Getenv("pgUser"), os.Getenv("pgPassword"), os.Getenv("pgDbname"))
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	url := context.PostForm("url")
	shortUrl := GetRandomStr(8)
	for i := 0; i < 3; i++ {
		insertSql := `INSERT INTO "url-shortener" ("short", "origin") values ($1, $2)`
		_, err = db.Exec(insertSql, shortUrl, url)
		if err != nil {
			context.HTML(http.StatusOK, "index.html", gin.H{
				"error": err,
			})
			fmt.Println(err)
			return
		} else {
			break
		}
	}

	context.HTML(http.StatusOK, "index.html", gin.H{
		"shortUrl": shortUrl,
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
