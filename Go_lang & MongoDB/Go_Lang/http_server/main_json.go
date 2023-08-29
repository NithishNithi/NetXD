package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Article struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Summary string `json:"Summary"`
	Author  string `json:"Author"`
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/product", product)
	http.HandleFunc("/article", ArticleHandler)
	http.HandleFunc("/getarticle",getarticle)
	fmt.Println("Running on server : 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Fprintf(w, "<h1>Welcome to home</h1>")
}

func product(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>You are in the Product Section</h1>")
}

func ArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		reqBody, _ := io.ReadAll(req.Body)
		var post Article
		err := json.Unmarshal(reqBody, &post)
		post.Author = "jhon"
		if err != nil {
			fmt.Println("got an error")
			fmt.Fprintf(w, err.Error())

		} else {
			json.NewEncoder(w).Encode(post)
		}
	} else {
		w.WriteHeader((http.StatusBadRequest))
		fmt.Fprintf(w, "<h1>UNABLE TO PROCESS YOUR REQUEST</h1>")

	}

}
func getarticle(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		reqbody,_:= io.ReadAll(req.Body)
		var articles [] Article
		err := json.Unmarshal(reqbody, &articles)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w,err.Error())
		} else{
			fmt.Println(articles)
			articles = append(articles, Article{ID:"3",Title: "MIB"})
			json.NewEncoder(w).Encode(articles)
		}
	}
}
