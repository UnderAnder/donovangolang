// Lissajous генерирует анимированный GIF из случайных
// фигур Лиссажу.
package lissajous

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 255, 255, 255}, color.RGBA{255, 255, 0, 255}, color.RGBA{0, 0, 255, 255}}

func Lissajous(out io.Writer, cycles, size, nframes, delay int) {
	const (
		//cycles  = 7      // number of complete x oscillator revolutions
		res = 0.0001 // angular resolution
		//size    = 100    // image canvas covers [-size..+size]
		//nframes = 32     // number of animation frames
		//delay   = 4      // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			colorIndex := uint8(rand.Intn(4) + 1)
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	}
}
