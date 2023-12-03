package controller

import (
	"fmt"
	// "hello/go/Go-projects/blog-rest-api/controller"
	// controller "hello/go/Go-projects/blog-rest-api/controller/blog"
	"github.com/karisa/Go-projects/blog-rest-api/controller"
	"github.com/karisa/Go-projects/blog-rest-api/controller/blog"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandler() {
	router.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")
	router.HandleFunc("/api/post/{id}", controller.GetPost).Methods("GET")
	router.HandleFunc("/api/post/new", controller.CreatePost).Methods("POST")
	router.HandleFunc("/api/post/update", controller.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/post/delet/{id}", controller.DeletePost).Methods("DELETE")
}

func Start() {
	router = mux.NewRouter()

	initHandler()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}