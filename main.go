package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello, world! (from a VM)")
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
