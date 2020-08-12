// Package surface вычисляет SVG-представление трехмерного графика функции.
package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 1200, 640           // размер канвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	xyrange       = 30.0                // Диапазоп осей
	xyscale       = width / 2 / xyrange // Пикселей в единице x или z
	zscale        = height * 0.4        // Пикселей в единице z
	angle         = math.Pi / 6         // Углы осей x, y (=30 градусам)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin 30, cos 30
// Surface вычисляет SVG-представление трехмерного графика функции.
func Surface(w io.Writer, width int, height int, color string) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, z := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			if z >= 0 {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='rgb(%2.1f%%,0%%,0%%)'/>\n", //#%02x00%02x rgb(%02d%%,00%%,%02d%%)
					ax, ay, bx, by, cx, cy, dx, dy, z*100)
			} else {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='rgb(0%%,0%%,%2.1f%%)'/>\n", //#%02x00%02x rgb(%02d%%,00%%,%02d%%)
					ax, ay, bx, by, cx, cy, dx, dy, z*-100)
			}

		}
		//fmt.Fprintf(w, "<linearGradient id='g'><stop offset='0' stop-color='red'/><stop offset='1' stop-color='blue'/></linearGradient>")
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Ищем угловую точку (x,y) ячейки (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

/* TODO: Упражнение 3.1. Если функция f возвращает значение float64, не являющееся
конечным, SVG-файл содержит неверные элементы <polygon> (хотя многие визуа-
лизаторы SVG успешно обрабатывают эту ситуацию). Измените программу так, что­
бы некорректные многоугольники были опущены.
Упражнение 3.2. Поэкспериментируйте с визуализациями других функций из па­
кета math. Сможете ли вы получить изображения наподобие коробки для яиц, седла
или холма?
Упражнение 3.3. Окрасьте каждый многоугольник цветом, зависящим от его вы­
соты, так, чтобы пики были красными (#ff0000), а низины — синими (#0000ff).
Упражнение 3.4. Следуя подходу, использованному в примере с фигурами Лисса­
жу из раздела 1.7, создайте веб-сервер, который вычисляет поверхности и возвращает
клиенту SVG-данные. Сервер должен использовать в ответе заголовок ContentType
наподобие следующего:
w.Header().Set("ContentType", "image/svg+xml")
(Этот шаг не был нужен в примере с фигурами Лиссажу, так как сервер использует
стандартную эвристику распознавания распространенных форматов наподобие PNG
по первым 512 байтам ответа и генерирует корректный заголовок.) Позвольте клиен­
ту указывать разные параметры, такие как высота, ширина и цвет, в запросе HTTP. */
