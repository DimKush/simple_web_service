package process

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ResponceSize(url string) error {
	fmt.Println("Getting responce", url)
	responce, err := http.Get(url)

	if err != nil {
		return err
	}

	defer responce.Body.Close()

	body, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		return err
	}

	fmt.Println(body)

	return nil
}
