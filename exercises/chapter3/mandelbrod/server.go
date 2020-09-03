package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/mandelbroad", ploter)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func ploter(w http.ResponseWriter, r *http.Request) {
	width, height := 1024, 1024
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if len(r.Form["width"]) != 0 {
		width, _ = strconv.Atoi(r.Form["width"][0])
	}

	if len(r.Form["height"]) != 0 {
		height, _ = strconv.Atoi(r.Form["height"][0])
	}
	Draw(w, width, height)
}
