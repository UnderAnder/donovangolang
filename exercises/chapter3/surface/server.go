package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/svg", ploter)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func ploter(w http.ResponseWriter, r *http.Request) {
	width, height := 640, 420
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	if len(r.Form["width"]) != 0 {
		width, _ = strconv.Atoi(r.Form["width"][0])
	}

	if len(r.Form["height"]) != 0 {
		height, _ = strconv.Atoi(r.Form["height"][0])
	}
	plotSVG(w, width, height)
}
