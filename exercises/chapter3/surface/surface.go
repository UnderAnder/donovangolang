package main

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func plotSVG(w io.Writer, width, height int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			//ax, ay := corner(i+1, j)
			//bx, by := corner(i, j)
			//cx, cy := corner(i, j+1)
			//dx, dy := corner(i+1, j+1)
			//fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
			//	ax, ay, bx, by, cx, cy, dx, dy)
			ax, ay, _ := corner(i+1, j)
			bx, by, z := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			if z >= 0 {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
					"fill='rgb(%2.1f%%,0%%,0%%)'/>\n", //#%02x00%02x rgb(%02d%%,00%%,%02d%%)
					ax, ay, bx, by, cx, cy, dx, dy, z*100)
			} else {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
					"fill='rgb(0%%,0%%,%2.1f%%)'/>\n", //#%02x00%02x rgb(%02d%%,00%%,%02d%%)
					ax, ay, bx, by, cx, cy, dx, dy, z*-100)
			}
		}
	}
	fmt.Fprint(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := safeValue(saddle(x, y))
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 23.0
	b := 15.0

	return (y*y)/(a*a) - (x*x)/(b*b)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func safeValue(v float64) float64 {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		v = 0
	}
	return v
}
