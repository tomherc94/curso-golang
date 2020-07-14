package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Google I/O 2012 - Go Concurrency Patterns

//<-chan - canal somente leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) { //funcao anonima
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) //aplica a funcao anonima (de forma goroutine) a todos urls
	}
	return c //é retornada antes da execução da requisição http da linha 17 devido a concorrencia
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
