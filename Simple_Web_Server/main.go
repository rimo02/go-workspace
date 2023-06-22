package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, " 404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, " method not suported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "HELLO")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "Post Request sent successfully\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Startign server at localhost 5050\n")
	if err := http.ListenAndServe(":5050", nil); err != nil {
		log.Fatal(err)
	}
}
