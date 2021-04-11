package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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

func creatHandle(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signature.txt", options, os.FileMode(0600))
	CatchError(err)
	_, err = fmt.Fprintln(file, signature)
	CatchError(err)
	err = file.Close()
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", IndexHandler)
	http.HandleFunc("/guestbook/new", NewHandler)
	http.HandleFunc("/guestbook/create", creatHandle)
	err := http.ListenAndServe("localhost:9100", nil)
	if err != nil {
		log.Fatal("Error", err)
	}
}
