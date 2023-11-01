package main

import (
	"github.com/karisa/golang-gin-poc/entity"
	"github.com/karisa/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	VideoController service.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func (ctx *gin.Context)  {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func (ctx *gin.Context)  {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}