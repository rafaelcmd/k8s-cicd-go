package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func raiz(num float64) string {
	x := num
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}

	return "Code.education Rocks!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, raiz(0.001))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
