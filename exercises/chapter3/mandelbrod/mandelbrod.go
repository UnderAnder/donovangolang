package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
	sswidth, ssheight      = width * 2, height * 2
	iterations             = 200
	red                    = 225
	green                  = 240
	blue                   = 250
)

var ssColors [sswidth][ssheight]color.Color

func Draw(w io.Writer, height, width int) {
	superSampling("newton")
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			si, sj := 2*px, 2*py

			r1, g1, b1, a1 := ssColors[si][sj].RGBA()
			r2, g2, b2, a2 := ssColors[si+1][sj].RGBA()
			r3, g3, b3, a3 := ssColors[si+1][sj+1].RGBA()
			r4, g4, b4, a4 := ssColors[si][sj+1].RGBA()

			avgColor := color.RGBA{
				uint8((r1 + r2 + r3 + r4) / 1028),
				uint8((g1 + g2 + g3 + g4) / 1028),
				uint8((b1 + b2 + b3 + b4) / 1028),
				uint8((a1 + a2 + a3 + a4) / 1028)}

			img.Set(px, py, avgColor)
		}
	}

	png.Encode(w, img)
}

func superSampling(f string) {
	for py := 0; py < ssheight; py++ {
		y := float64(py)/ssheight*(ymax-ymin) + ymin
		for px := 0; px < sswidth; px++ {
			x := float64(px)/sswidth*(xmax-xmin) + xmin
			var m int
			switch f {
			case "newton":
				m = newton(complex(x, y))
			default:
				m = mandelbroad(complex(x, y))
			}

			colour := color.NRGBA{R: uint8(red * m),
				G: uint8(green * m), B: uint8(blue * m), A: 255}
			ssColors[px][py] = colour
		}
	}
}

func mandelbroad(c complex128) int {
	i := 0
	for z := c; cmplx.Abs(z) <= 2 && i < iterations; i++ {
		z = z*z + c
	}
	return i
}

func newton(z complex128) int {
	for i := 0; i < iterations; i++ {
		z -= (z - 1/(cmplx.Pow(z, 3))) / 4
		if cmplx.Abs(cmplx.Pow(z, 4)-1) < 1e-6 {
			return i
		}
	}
	return 0
}
