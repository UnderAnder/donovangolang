// Dupl выводит текст каждой строки, которая появляется в
// стандартном вводе более одного раза, а также количество
// ее появлений,
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // Из цикла не выходит?
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	fmt.Println("1 цикл отработал")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
