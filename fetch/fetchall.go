// Fetchall выполняет параллельную выборку URL и сообщает
// о затраченном времени и размере ответа для каждого из них.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	var uri string
	var nbytes int64
	if !strings.HasPrefix(url, "http://") {
		uri = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err = io.Copy(ioutil.Discard, resp.Body)
	aa, err := ioutil.ReadAll(resp.Body)
	if err := ioutil.WriteFile("./fetch/"+url, aa, 0644); err != nil {
		ch <- fmt.Sprintf("While writing file %s: %v	", url, err)
		return
	}
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v	", uri, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, uri)
}

/*Упражнение 1.10. Найдите веб-сайт, который содержит большое количество дан­
ных. Исследуйте работу кеширования путем двукратного запуска fe tc h a ll и срав­
нения времени запросов. Получаете ли вы каждый раз одно и то же содержимое?
Измените fe tc h a ll так, чтобы вывод осуществлялся в файл и чтобы затем можно
было его изучить.
Упражнение 1.11. Выполните f e tc h a ll с длинным списком аргументов, таким
как образцы, доступные на сайте a le x a . com. Как ведет себя программа, когда веб­
сайт просто не отвечает? (В разделе 8.9 описан механизм отслеживания таких си­
туаций.)*/
