package main

import (
	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", handler)
	r.GET("/count", counter)
	r.Run(":9999")
}

func handler(c *gee.Context) {
	html := "<h2>hello</h2>"
	c.HTML(200, html)
}

func counter(c *gee.Context) {
	c.Data(200, []byte("fadsfsa"))
}
