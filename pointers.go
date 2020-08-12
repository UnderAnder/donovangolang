package main

import "fmt"

func main() {
	x := 1
	p := &x         // р имеет тип *int и указывает на х
	fmt.Println(*p) // "1"
	*p = 44         // Эквивалентно присваиванию х = 2
	fmt.Println(x)  // "2"

	р := f()
	fmt.Println(р)
}

func f() *int {
	v := 1
	return &v
}
