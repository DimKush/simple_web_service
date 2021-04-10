package main

import (
	"html/template"
	"log"
	"net/http"
	"simple_web_service/internal/datafile"
	guestboook "simple_web_service/internal/guestbook"
)

func CatchError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	signature, err := datafile.GetStringsFromFile("signature.txt")
	if err != nil {
		log.Fatal(err)
	}

	html, err := template.ParseFiles("index.html")
	CatchError(err)

	guestboook := guestboook.Guestbook{
		SignatureCount: len(signature),
		Signature:      signature,
	}

	err = html.Execute(writer, guestboook)
	CatchError(err)
}

func NewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	CatchError(err)

	signature, err := datafile.GetStringsFromFile("signature.txt")
	CatchError(err)

	guestboook := guestboook.Guestbook{
		SignatureCount: len(signature),
		Signature:      signature,
	}

	err = html.Execute(writer, guestboook)
	CatchError(err)

}

func main() {
	http.HandleFunc("/guestbook", IndexHandler)
	http.HandleFunc("/guestbook/new", NewHandler)

	err := http.ListenAndServe("localhost:9100", nil)
	if err != nil {
		log.Fatal("Error", err)
	}
}
