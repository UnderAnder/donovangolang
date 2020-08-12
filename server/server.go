// Server2 - минимальный "echo"-сервер со счетчиком запросов,
package main

import (
	"fmt"
	"learninggo/mandelbrot"
	"learninggo/surface"
	"lissajous"
	"log"
	"net/http"
	"strings"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", indexer)
	http.HandleFunc("/gif", gifer)
	http.HandleFunc("/svg", svger)
	http.HandleFunc("/png", pnger)
	http.HandleFunc("/info", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

// Обработчик, возвращающий компонент пути запрашиваемого URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func indexer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/info'><span>info</span></a>\n")
	fmt.Fprintf(w, "<a href='/count'><span>count</span></a>\n")
	fmt.Fprintf(w, "<a href='/gif'><span>gif</span></a>\n")
	fmt.Fprintf(w, "<a href='/svg'><span>svg</span></a>\n")
	fmt.Fprintf(w, "<a href='/png'><span>Mandelbrot</span></a>\n")
	fmt.Fprintf(w, "<a href='/gif'><span>count</span></a>\n")
}

// Счетчик, возвращающий количество сделанных запросов,
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func gifer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	//fmt.Println(r.Form["cycles"])
	cycles := 5
	size := 20
	//cycles, _ = strconv.Atoi(arrToString(r.Form["cycles"]))
	//size, _ = strconv.Atoi(arrToString(r.Form["size"]))
	lissajous.Lissajous(w, cycles, size, 32, 7)
	//cycles  = 7      // number of complete x oscillator revolutions
	//size    = 100    // image canvas covers [-size..+size]
	//nframes = 32     // number of animation frames
	//delay   = 4      // delay between frames in 10ms units
}

func svger(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	width := 1200
	height := 640
	color := "white"
	//width, _ = strconv.Atoi(arrToString(r.Form["width"]))
	//height, _ = strconv.Atoi(arrToString(r.Form["height"]))
	surface.Surface(w, width, height, color)
}

func pnger(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	mandelbrot.Draw(w)
}

func arrToString(strArray []string) string {
	return strings.Join(strArray, " ")
}

/*Упражнение 1.12. Измените сервер с фигурами Лиссажу так, чтобы значения
параметров считывались из URL. Например, URL вида h ttp ://lo c a lh o s t: 8000/
?cycles=20 устанавливает количество циклов равным 20 вместо значения по умол­
чанию, равного 5. Используйте функцию strco n v .A to i для преобразования строко­
вого параметра в целое число. Просмотреть документацию по данной функции мож­
но с помощью команды go doc strco n v .A to i.*/
