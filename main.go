package main

import (
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/videos", handlers.GetVideos)
	r.GET("/create", handlers.CreateVideo)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
