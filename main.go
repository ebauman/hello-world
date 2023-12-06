package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	htmlTemplate "html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var t, err = htmlTemplate.New("template").Parse(Template)
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte("error parsing template: " + err.Error()))
			return
		}

		var strings = []string{}
		strings = append(strings, os.Environ()...)

		var out = struct {
			Output []string
		}{
			Output: strings,
		}

		t.Execute(writer, out)

		fmt.Printf("REQUEST! %s, %s, %s\n", time.Now(), request.RemoteAddr, request.Method)
	})

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
