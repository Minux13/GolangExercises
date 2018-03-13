package main

import (
	"github.com/gorilla/mux"
	//"fmt"
	"net/http"
	"time"
	"encoding/json"
	"log"
	"strconv"
)

type Note struct{
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}

var noteStore = make(map[string]Note)

var id int

func GetNote(w http.ResponseWriter, r *http.Request){
	var notes []Note
	for _, v := range noteStore{
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil{
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func PostNote(w http.ResponseWriter, r *http.Request){
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil{
		panic(err)
	}
	note.CreatedAt = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(note)
	if err != nil{
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/GetNote", GetNote).Methods("GET")
	r.HandleFunc("/PostNote", PostNote).Methods("POST")
	//r.HandleFunc("/PutNote/{id}", PutNote).Methods("PUT")
	//r.HandleFunc("/DeleteNote{id}", DeleteNote).Methods("Delete")

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

