package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	sample     = "/image"
	samplePath = "./image.jpg"
	index      = "/"
)

func hello(w http.ResponseWriter, req *http.Request) {
	//fmt.Println(w, req.Body)
	http.ServeFile(w, req, samplePath)
}

func main() {
	fmt.Println("Starting server on 8064")
	http.HandleFunc(sample, hello)
	http.HandleFunc(index, func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Starting Page")
	})

	log.Fatal(http.ListenAndServe(":8064", nil))
}
