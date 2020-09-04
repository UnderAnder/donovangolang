package main

import (
	"./mandelbrod"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	const addr = "localhost:8001"
	fmt.Printf("Start server at %v\n", addr)
	fmt.Printf("Mandelbroad http://%v/fractal\n", addr)
	fmt.Printf("Newton http://%v/fractal?f=newton&width=640&height=640\n", addr)
	http.HandleFunc("/fractal", plotter)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func plotter(w http.ResponseWriter, r *http.Request) {
	width, height := 1024, 1024
	f := "mandelbrod"
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if len(r.Form["width"]) != 0 {
		width, _ = strconv.Atoi(r.Form["width"][0])
	}

	if len(r.Form["height"]) != 0 {
		height, _ = strconv.Atoi(r.Form["height"][0])
	}
	if len(r.Form["f"]) != 0 {
		f = r.Form["f"][0]
	}
	mandelbrod.Draw(w, width, height, f)
}
