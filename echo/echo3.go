// Echo3 выводит аргументы командной строки
package main

import (
	"fmt"
	"os"
)

func main() {
	//	fmt.Println(strings.Join(os.Args[0:], " "))
	//	fmt.Println("---")
	fmt.Println(os.Args[0])

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
		//s += sep + arg
		//sep = " "
	}
}

/* TODO: Упражнение 1.1. Измените программу echo так, чтобы она выводила также
o s.A rg s[0], имя выполняемой команды.
Упражнение 1.2. Измените программу echo так, чтобы она выводила индекс и
значение каждого аргумента по одному аргументу в строке.
Упражнение 1.3. Поэкспериментируйте с измерением разницы времени выполне­
ния потенциально неэффективных версий и версии с применением s trin g s . Doin.
(В разделе 1.6 демонстрируется часть пакета tim e, а в разделе 11.4 — как написать
тест производительности для ее систематической оценки.)*/
