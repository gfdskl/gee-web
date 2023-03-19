package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", handler)
	r.GET("/count", counter)
	r.Run(":9999")
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world\n"))
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", r.URL.Path)
}
