package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello<h1>")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Start server....")
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye!"))
}
