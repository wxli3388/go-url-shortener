package main

import (
	"net/http"
	"url-shortener/interal/wire"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.LoadHTMLGlob("./../web/*")

	urlController := wire.InitializeUrlController()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/list", urlController.List)
	router.GET("/:shortUrl", urlController.Redirect)
	router.POST("/generate", urlController.Generate)

	router.Run(":8080")
}
