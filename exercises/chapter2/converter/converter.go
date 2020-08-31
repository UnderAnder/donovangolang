package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	if len(os.Args) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			t, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "converter: %v\n", err)
				os.Exit(1)
			}
			convert(t)
		}
		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
			os.Exit(1)
		}
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "converter: %v\n", err)
				os.Exit(1)
			}
			convert(t)
		}
	}

}

func convert(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	m := Metre(t)
	foot := Foot(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Printf("%s = %s, %s = %s\n", m, MToF(m), foot, FToM(foot))
}
