package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Post structure
type Post struct {
	Title   string `json:"title"`
	Desc    string `json:"desk"`
	Content string `json:"content"`
}

// Posts is array of Post
type Posts []Post

func allPosts(w http.ResponseWriter, r *http.Request) {
	posts := Posts{
		Post{
			Title:   "test title",
			Desc:    "test description",
			Content: "Hello world",
		},
	}

	fmt.Println("Endpoint: All posts Endpoint")
	json.NewEncoder(w).Encode(posts)
}

func testPostPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint Hit")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/posts", allPosts).Methods("GET")
	router.HandleFunc("/posts", testPostPosts).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}
