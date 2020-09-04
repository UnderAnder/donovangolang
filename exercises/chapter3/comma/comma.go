package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

func comma(s string) string {
	var buf bytes.Buffer
	i := (3 - utf8.RuneCountInString(s)%3) % 3
	for _, r := range s {
		if i == 3 {
			buf.WriteByte(',')
			i = 0
		}
		buf.WriteRune(r)
		i++
	}

	return buf.String()
}

func commaSign(s string) string {
	start, stop := 0, len(s)
	if strings.HasPrefix(s, "-") {
		start = 1
	}
	if strings.ContainsRune(s, '.') {
		stop = strings.IndexRune(s, '.')
	}
	return s[:start] + comma(s[start:stop]) + s[stop:]
}

func main() {
	fmt.Println(commaRecursive("12345.6789123"))
	fmt.Println(commaRecursive("-100000000000"))
	fmt.Println(commaRecursive("gearbox"))
	fmt.Println(comma("12345.6789123"))
	fmt.Println(comma("-100000000000"))
	fmt.Println(comma("gearbox"))
	fmt.Println(commaSign("12345.6789123"))
	fmt.Println(commaSign("-100000000000"))
	fmt.Println(commaSign("gearbox"))
}
