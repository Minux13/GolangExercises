package main

import (
    "fmt"
    "net/http"
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
 	servermux := http.NewServeMux()
	servermux.HandleFunc("/", index)
	servermux.HandleFunc("/dos", dos)
	servermux.HandleFunc("/tres", tres)

	http.ListenAndServe(":8080", servermux)
}

