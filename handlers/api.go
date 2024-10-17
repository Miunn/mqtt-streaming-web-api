package handlers

import (
	"main/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVideos(c *gin.Context) {
	db := database.GetDatabase()
	var videos []database.Video
	db.Find(&videos)
	c.JSON(http.StatusOK, videos)
}

func CreateVideo(c *gin.Context) {
	db := database.GetDatabase()
	video := database.Video{VideoID: c.Query("video_id")}
	db.Create(&video)
	c.JSON(http.StatusOK, video)
}
