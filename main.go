package main

import (
	"fmt"
	"learninggo/comma"
	popcount "learninggo/popcount4"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("x\t %#08b\t %#[1]d\n", x)
	fmt.Printf("y\t %#08b\t %#[1]d\n", y)
	fmt.Printf("x&y\t %#08b\t %#[1]d\n", x&y)
	fmt.Printf("x|y\t %#08b\t %#[1]d\n", x|y)
	fmt.Printf("x^y\t %#08b\t %#[1]d\n", x^y)
	fmt.Printf("x&^y\t %#08b\t %#[1]d\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // Проверка принадлежности множеству
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) //"01000100", множество {2,6}
	fmt.Printf("%08b\n", x>>1) //"00010001", множество {0,4}
	fmt.Println(popcount.PopCount(110))

	ascii := 'а'
	Unicode := '★'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a ’a'"
	fmt.Printf("%d %[1]c %[1]q\n", Unicode) // "9733 ★
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n,M

	fmt.Printf(comma.Comma("123456789"))
}
