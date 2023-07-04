package main

import (
	htmlTemplate "html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var t = htmlTemplate.Must(htmlTemplate.ParseFiles("./html.tmpl"))
		var strings = []string{}
		strings = append(strings, os.Environ()...)

		var out = struct {
			Output []string
		}{
			Output: strings,
		}

		t.Execute(writer, out)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
