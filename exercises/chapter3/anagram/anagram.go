package main

import "fmt"

func anagram(a, b string) bool {
	return eq(runeOccur(a), runeOccur(b))
}

func runeOccur(s string) map[rune]int {
	occur := make(map[rune]int)
	for _, r := range s {
		occur[r]++
	}
	return occur
}

func eq(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}
func main() {
	s1 := "молоко"
	s2 := "локомо"
	s3 := "около"
	s4 := "молока"
	fmt.Println(anagram(s1, s1))
	fmt.Println(anagram(s1, s2))
	fmt.Println(anagram(s1, s3))
	fmt.Println(anagram(s1, s4))
}
