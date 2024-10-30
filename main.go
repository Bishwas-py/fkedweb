package main

import (
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	Name        string `json:"name"`
	CompletedAt string `json:"completed_at"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fprintf, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	if err != nil {
		return
	}
	fmt.Println(fprintf)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
