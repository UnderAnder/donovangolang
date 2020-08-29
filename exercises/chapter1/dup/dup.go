//Dup2 выводит текст каждой строки, которая появляется во
// входных данных более одного раза. Программа читает
// стандартный ввод или список именованных файлов

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		//Упражнение 1.4. Измените программу dup2 так, чтобы она выводила имена всех
		//файлов, в которых найдены повторяющиеся строки
		counts[f.Name()+"\t"+input.Text()]++
	}
}
