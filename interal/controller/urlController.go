package controller

import (
	"html/template"
	"net/http"
	"url-shortener/interal/service"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	urlService service.UrlService
}

func NewUrlController(urlService service.UrlService) UrlController {
	return UrlController{
		urlService: urlService,
	}
}

func (controller *UrlController) Generate(context *gin.Context) {
	shortUrl, err := controller.urlService.Generate(context)
	if err != nil {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"error": err,
		})
	}
	qrcode, err := controller.urlService.GetQRcode(context.Request.Host, shortUrl)
	context.HTML(http.StatusOK, "index.html", gin.H{
		"shortUrl": shortUrl,
		"qrcode":   template.URL(qrcode),
		"error":    "",
	})
}

func (controller *UrlController) List(context *gin.Context) {
	shortUrls, err := controller.urlService.GetShortUrl(context)
	if err != nil {
		context.HTML(http.StatusOK, "list.html", gin.H{
			"error": "Something went wrong...",
		})
	}
	context.HTML(http.StatusOK, "list.html", gin.H{
		"error": "",
		"data":  shortUrls,
	})
}

func (controller *UrlController) Redirect(context *gin.Context) {
	url, err := controller.urlService.GetOriginUrl(context)
	if err != nil {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"error": "Failed to connect to the database",
		})
	}
	context.Redirect(http.StatusMovedPermanently, url)
}
