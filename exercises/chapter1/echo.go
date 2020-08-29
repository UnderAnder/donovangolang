// Echo1 выводит аргументы командной строки
/* TODO:
Упражнение 1.1. Измените программу echo так, чтобы она выводила также
os.Args[0], имя выполняемой команды.
Упражнение 1.2. Измените программу echo так, чтобы она выводила индекс и
значение каждого аргумента по одному аргументу в строке.
Упражнение 1.3. Поэкспериментируйте с измерением разницы времени выполне
ния потенциально неэффективных версий и версии с применением strings.Join.
(В разделе 1.6 демонстрируется часть пакета tim e, а в разделе 11.4 — как написать
тест производительности для ее систематической оценки.)*/
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(os.Args[0])

	printArgsRange(start)
	printArgsFor(start)

}

func printArgsRange(start time.Time) {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
	fmt.Printf("%v микросекунд прошло\n", time.Since(start).Microseconds())
}

func printArgsFor(start time.Time) {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

}
