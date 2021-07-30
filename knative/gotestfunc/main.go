package main

import (
	"fmt"
	"math"
	"math/rand"
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

func main() {
	fmt.Print(load())
	return
}
