package main

import (
	"github.com/gorilla/mux"
	//"fmt"
	"net/http"
	"time"
	"log"
	"strconv"
	"encoding/json"
)
type Note struct{
	Titulo string `json: "title"`
	Description string `json:description"`
	CreatedAt time.Time `json:"created_at"`
}

var noteStore = make(map[string] Note)

var id int

func dos(w http.ResponseWriter, r *http.Request){
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil{
		panic(err)
	}
	w.WriteHeader(200)
	w.Write(j)
}

func tres(w http.ResponseWriter, r *http.Request){
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil{
		panic(err)
	}
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	w.WriteHeader(200)
	//w.Write(k)
	//fmt.Fprintf(w, "Dooos")		

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

