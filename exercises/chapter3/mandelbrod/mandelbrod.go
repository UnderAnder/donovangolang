package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
	iterations             = 200
	red                    = 225
	green                  = 240
	blue                   = 250
)

func Draw(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			m := mandelbroad(complex(x, y))

			col := color.NRGBA{R: uint8(red * m),
				G: uint8(green * m), B: uint8(blue * m), A: 255}
			img.Set(px, py, col)
		}
	}
	png.Encode(w, img)
}

func mandelbroad(c complex128) int {
	i := 0
	for z := c; cmplx.Abs(z) <= 2 && i < iterations; i++ {
		z = z*z + c
	}
	return i
}

func main() {
	Draw(os.Stdout)
}
