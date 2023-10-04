package main
 import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
 )

 type Movie struct{
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json: "director"`

 }

 type Director struct{
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`

 }

 var movies []Movie

 func getMovies(w, http.ResponseWriter, r *http.Request)

 func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "48322", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})


	r.HandleFunc("/movies", getMovies).Method("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")

	fmt.Printf("Starting the server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

 }