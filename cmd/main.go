package main

import (
	"fmt"
	process "simple_web_service/internal"
)

func main() {
	chnPckg := make(chan process.Page)
	go process.ResponceSize("http://google.com", chnPckg)
	go process.ResponceSize("http://golang.org", chnPckg)

	//var pg process.Page
	pg := <-chnPckg
	fmt.Println(pg.GetUrl())
	fmt.Println(pg.GetBodySize())
	pg = <-chnPckg
	fmt.Println(pg.GetUrl())
	fmt.Println(pg.GetBodySize())
}
