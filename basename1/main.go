// basename убирает компоненты каталога и суффикс типа файла.
// а => a, a.go => a, a/b/c.go => с, a/b.c.go => b.c
package basename

import "strings"

func basename(s string) string {
	// Отбрасываем последний символ '/' и все перед ним.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Сохраняем все до последней точки '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[dot:]
	}
	return s
}
