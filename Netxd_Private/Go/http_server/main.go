package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHello)
	http.HandleFunc("/product", product)
	http.HandleFunc("/contact", contact)
	log.Fatal(http.ListenAndServe(":2000", nil))
	fmt.Println("running through port no 2000")
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<h1>PageNotFound</h1>")
	}else{
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Fprintf(w, "<h1>Welcome to Http </h1>")
	}
	// 404 - page not found statusNot found
	// 403 - Forbidden
	// 500 - Internal server error
}
func product(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Fprintf(w, "<h1>Welcome to product section</h1>")
}

func contact(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Fprintf(w, "<h1>Mr Nithish T</h1>")
}
