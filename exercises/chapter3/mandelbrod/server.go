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
	fmt.Printf("Newton http://%v/fractal?f=newton&x=0.11&y=0.3&zoom=1.5\n", addr)
	http.HandleFunc("/fractal", plotter)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func plotter(w http.ResponseWriter, r *http.Request) {
	x, y, zoom := 0.0, 0.0, 0.0
	f := "mandelbrod"
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if len(r.Form["x"]) != 0 {
		x, _ = strconv.ParseFloat(r.Form["x"][0], 64)
	}

	if len(r.Form["y"]) != 0 {
		y, _ = strconv.ParseFloat(r.Form["y"][0], 64)
	}
	if len(r.Form["f"]) != 0 {
		f = r.Form["f"][0]
	}
	if len(r.Form["zoom"]) != 0 {
		zoom, _ = strconv.ParseFloat(r.Form["zoom"][0], 64)
	}
	mandelbrod.Draw(w, x, y, zoom, f)
}
