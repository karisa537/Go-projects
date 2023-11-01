package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/karisa/golang-gin-poc/entity"
	"github.com/karisa/golang-gin-poc/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct{
	service service.VideoService
}

func New(service service.VideoController) VideoController {
	return controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}