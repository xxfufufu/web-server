package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	file_server := http.FileServer(http.Dir("./static"))
	http.Handle("/", file_server)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/form", handlerForm)

	fmt.Println("running on port: 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "hello!")
}

func handlerForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name :%s\n", name)
	fmt.Fprintf(w, "address :%s\n", address)
}
