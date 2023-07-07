package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
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

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9000", nil)

	log.Fatal(http.ListenAndServe(":80", nil))
}
