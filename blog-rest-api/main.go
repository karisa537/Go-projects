package main

import (
	"github.com/karisa/Go-projects/blog-rest-api/model"
	"github.com/karisa/Go-projects/blog-rest-api/controller"
	// "hello/go/Go-projects/blog-rest-api/controller"
	// "hello/go/Go-projects/blog-rest-api/model"
)

func main() {
	model.Init()
	controller.Start()
}