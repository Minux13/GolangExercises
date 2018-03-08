package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"time"
	"log"
)

func dos(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Uno")		
}
func tres(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Dooos")		
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/dos", dos).Methods("GET")
	r.HandleFunc("/tres", tres).Methods("POST")

	server := &http.Server{
			Addr			: ":8080",
			Handler			: r,
			ReadTimeout		: 10 * time.Second,
			WriteTimeout	: 10 * time.Second,
			MaxHeaderBytes	: 1 << 20,
	}

	log.Println("Sirviendo....")

	server.ListenAndServe()
}

