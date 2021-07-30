package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"
)

const param = 100000

func load() (res float64) {
	var a [param]float64
	start := rand.Float64()
	for _, d := range a {
		d = math.Sin(start)
		start = d
	}
	res = start
	return res
}

func welcome(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	res := load()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Fprintf(w, "%f %s", res, elapsed.String())
}

func main() {
	http.HandleFunc("/", welcome)
	http.ListenAndServe(":8080", nil)
	return
}
