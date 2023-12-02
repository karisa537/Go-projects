package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){

	articles := Articles{
		Article{Title: "Test file", Desc: "Test Description", Content: "Hello karisa"},
	}
	fmt.Fprint(w, "All articles endpoint hit")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, " endpoint")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Homepage endpoint hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}