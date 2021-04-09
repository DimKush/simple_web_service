package process

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	url      string
	bodySize int
	body     string
}

// functions member
func (pg *Page) SetUrl(url string) {
	pg.url = url
}

func (pg *Page) SetBodySize(sz int) {
	pg.bodySize = sz
}

func (pg *Page) SetBody(body string) {
	pg.body = body
}

func (pg Page) GetUrl() string {
	return pg.url
}

func (pg Page) GetBodySize() int {
	return pg.bodySize
}

func (pg Page) GetBody() string {
	return pg.body
}

func ResponceSize(urlStr string, channel chan Page) {
	fmt.Println("Getting responce", urlStr)
	responce, err := http.Get(urlStr)

	if err != nil {
		log.Fatal(err)
	}

	defer responce.Body.Close()

	body, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- Page{url: urlStr, body: string(body), bodySize: len(body)}
}
