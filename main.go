package main

import (
	"fmt"
	"net/http"

	"github.com/abelardolugo/go-url-shortener.git/handler"
	"github.com/abelardolugo/go-url-shortener.git/store"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hey Go URL Shortener",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
