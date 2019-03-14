package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
        fmt.Fprintf(w, "Hello IOP 20180725")
	fmt.Println("helloHandler is called.")
}

func helloWhoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
    fmt.Fprintf(w, "Hello IOP Caleb")
	w.Write([]byte("Hello IOP Caleb"))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
    mux.HandleFunc("/hello-who", helloWhoHandler)
	http.ListenAndServe("0.0.0.0:8888", mux)
}