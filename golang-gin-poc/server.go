// package main

// import (
// 	// "github.com/karisa/golang-gin-poc/entity"
// 	"github.com/karisa/golang-gin-poc/service"
// 	"github.com/karisa/golang-gin-poc/controller"
// 	"github.com/gin-gonic/gin"
// )

// var (
// 	videoService service.VideoService = service.New()
// 	videoController service.VideoController = controller.New(videoService)
// )

// func main() {
// 	server := gin.Default()

// 	server.GET("/videos", func (ctx *gin.Context)  {
// 		ctx.JSON(200, videoController.FindAll())
// 	})

// 	server.POST("/videos", func (ctx *gin.Context)  {
// 		ctx.JSON(200, videoController.Save(ctx))
// 	})

// 	server.Run(":8080")
// }


package main

import (
	"github.com/karisa/golang-gin-poc/service"
	"github.com/karisa/golang-gin-poc/controller"
	"github.com/gin-gonic/gin"
	"github.com/karisa/golang-gin-poc/entity" // Add the import for entity
)

func main() {
	videoService := service.New()
	videoController := controller.New(videoService)

	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		var video entity.Video // Create a Video instance
		ctx.BindJSON(&video)
		videoController.Save(ctx, video) // Pass ctx and video to Save
		ctx.JSON(200, video)
	})

	server.Run(":8080")
}