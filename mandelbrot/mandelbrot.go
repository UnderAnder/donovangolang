package mandelbrot

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"sync"
	"time"
)

func Draw(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -1.25, .5, 1.25
		width, height          = 4096, 4096
		sswidth, ssheight      = width * 2, height * 2
	)
	var superColors [sswidth][ssheight]color.Color

	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(ssheight)
	for py := 0; py < ssheight; py++ {
		go func(py int) {
			defer wg.Done()
			y := float64(py)/ssheight*(ymax-ymin) + ymin
			for px := 0; px < sswidth; px++ {
				x := float64(px)/sswidth*(xmax-xmin) + xmin
				z := complex(x, y) // Точка (px, py) представляет комплексное значение z
				superColors[px][py] = newton(z)
				//img.Set(px, py, mandelbrot(z))
			}
		}(py)
	}
	wg.Wait()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			si, sj := 2*i, 2*j

			r1, g1, b1, a1 := superColors[si][sj].RGBA()
			r2, g2, b2, a2 := superColors[si+1][sj].RGBA()
			r3, g3, b3, a3 := superColors[si+1][sj+1].RGBA()
			r4, g4, b4, a4 := superColors[si][sj+1].RGBA()

			avgColor := color.RGBA{
				uint8((r1 + r2 + r3 + r4) / 1028),
				uint8((g1 + g2 + g3 + g4) / 1028),
				uint8((b1 + b2 + b3 + b4) / 1028),
				uint8((a1 + a2 + a3 + a4) / 1028)}

			img.Set(i, j, avgColor)
		}
	}

	fmt.Println(time.Since(start).Seconds())
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 5
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 4 {
			//return color.RGBA{255 - contrast*n, 255 - contrast*n, 255 - contrast*n, 255}
			green := uint8(255)
			blue := uint8(255) //uint8(imag(v)*128) + 127
			red := uint8(255 - contrast*n)
			return color.RGBA{red, green, blue, 255}
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

/*Упражнение 3.5. Реализуйте полноцветное множество Мандельброта с использо­
ванием функции image. NewRGBA и типа color. RGB А или color. YCbCr.
Упражнение 3.6. Супервыборка (supersampling) — это способ уменьшить эффект
пикселизации путем вычисления значений цвета в нескольких точках в пределах каж­
дого пикселя и их усреднения. Проще всего разделить каждый пиксель на четыре
“подпикселя”. Реализуйте описанный метод.
Упражнение 3.7. Еще один простой фрактал использует метод Ньютона для поис­
ка комплексных решений уравнения z*-\ = 0. Закрасьте каждую точку цветом, соот­
ветствующим тому корню из четырех, которого она достигает, а интенсивность цвета
должна соответствовать количеству итераций, необходимых для приближения к этому
корню.
Упражнение 3.8. Визуализация фракталов при высоком разрешении требует высо­
кой арифметической точности. Реализуйте один и тот же фрактал с помощью четырех
различных представлений чисел: complex64, complexl28, big.Float и big.Rat.
(Два последних типа вы найдете в пакете math/big. Float использует произволь­
ную, но ограниченную точность для чисел с плавающей точкой; Rat обеспечивает
неограниченную точность для рациональных чисел.) Сравните производительность и
потребление памяти при использовании разных типов. При каком уровне масштаби­
рования артефакты визуализации становятся видимыми?
Упражнение 3.9. Напишите программу веб-сервера, который визуализирует фрак­
талы и выводит данные изображения клиенту. Позвольте клиенту указывать значе­
ния х,у и масштабирования в качестве параметров HTTP-запроса.*/
