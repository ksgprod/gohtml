package html

import (
	"io"
	"net/http"
	"regexp"
)

// Titulo obtem o titulo de uma pagina html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			aRetorno := r.FindStringSubmatch(string(html))

			// teste para evitar erro
			if cap(aRetorno) == 0 {
				c <- "Erro ao ler a página " + url
				return
			}

			c <- aRetorno[1]
		}(url)
	}

	return c
}
