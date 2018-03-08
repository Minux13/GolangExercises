package main

import (
    "fmt"
    "net/http"
	"time"
	"log"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Principal</h1>")
}
func dos(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1 style=\"color: blue\" >Segunda</h1>")
}
func tres(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Tercera</h1>")
}

func main() {
 	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	mux.HandleFunc("/dos", dos)
	mux.HandleFunc("/tres", tres)

	server := &http.Server{
			Addr			: ":8080",
			Handler			: mux,
			ReadTimeout		: 10 * time.Second,
			WriteTimeout	: 10 * time.Second,
			MaxHeaderBytes	: 1 << 20,
	}

	log.Println("Sirviendo....")

	server.ListenAndServe()
}

