// Fetch выводит ответ на запрос по заданному URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

/*Упражнение 1.7. Вызов функции io .C o p y (d st,src ) выполняет чтение src и
запись в d st. Воспользуйтесь ею вместо io u til.R e a d A ll для копирования тела
ответа в поток o s.S td o u t без необходимости выделения достаточно большого для
хранения всего ответа буфера. Не забудьте проверить, не произошла ли ошибка при
вызове i o . Сору.
Упражнение 1.8. Измените программу fe tc h так, чтобы к каждому аргументу
URL автоматически добавлялся префикс h t t p :/ / в случае отсутствия в нем таково­
го. Можете воспользоваться функцией strin g s.H a sP re fix .
Упражнение 1.9. Измените программу fe tc h так, чтобы она выводила код состо­
яния HTTP, содержащийся в resp . S tatu s.*/
