package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	maxIterations          = 200
	red                    = 225
	green                  = 240
	blue                   = 250
)

func Draw(w io.Writer, height, width int, f string) {
	ssColors := superSampling(height, width, f)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			si, sj := 2*px, 2*py

			r1, g1, b1, a1 := ssColors[si][sj].RGBA()
			r2, g2, b2, a2 := ssColors[si+1][sj].RGBA()
			r3, g3, b3, a3 := ssColors[si+1][sj+1].RGBA()
			r4, g4, b4, a4 := ssColors[si][sj+1].RGBA()

			avgColor := color.RGBA{
				R: uint8((r1 + r2 + r3 + r4) / 1028),
				G: uint8((g1 + g2 + g3 + g4) / 1028),
				B: uint8((b1 + b2 + b3 + b4) / 1028),
				A: uint8((a1 + a2 + a3 + a4) / 1028)}

			img.Set(px, py, avgColor)
		}
	}

	err := png.Encode(w, img)
	if err != nil {
		fmt.Println(err)
	}
}

func superSampling(height, width int, f string) [][]color.Color {
	ssheight := height * 2
	sswidth := width * 2
	ssColors := make([][]color.Color, sswidth)
	for i := range ssColors {
		ssColors[i] = make([]color.Color, ssheight)
	}
	for py := 0; py < ssheight; py++ {
		y := float64(py)/float64(ssheight)*(ymax-ymin) + ymin
		for px := 0; px < sswidth; px++ {
			x := float64(px)/float64(sswidth)*(xmax-xmin) + xmin
			var m int
			switch f {
			case "newton":
				m = newton(complex(x, y))
			default:
				m = mandelbroad(complex(x, y))
			}

			colour := color.NRGBA{
				R: uint8(red * m),
				G: uint8(green * m),
				B: uint8(blue * m),
				A: 255}
			ssColors[px][py] = colour
		}
	}
	return ssColors
}

func mandelbroad(c complex128) int {
	i := 0
	for z := c; cmplx.Abs(z) <= 2 && i < maxIterations; i++ {
		z = z*z + c
	}
	return i
}

func newton(z complex128) int {
	for i := 0; i < maxIterations; i++ {
		z -= (z - 1/(cmplx.Pow(z, 3))) / 4
		if cmplx.Abs(cmplx.Pow(z, 4)-1) < 1e-6 {
			return i
		}
	}
	return 0
}
