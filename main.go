package main

import (
	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", handler)
	r.GET("/count", counter)
	r.GET("detail/:name", namer)

	// r.GET("/:age", func(c *gee.Context) {
	// 	c.String(200, c.Params["age"])
	// })
	// r.GET("age1", func(c *gee.Context) {
	// 	c.String(200, "age1")
	// })
	r.Run(":9999")
}

func namer(c *gee.Context) {
	obj := map[string]string{
		"name": c.Param("name"),
	}
	c.JSON(200, obj)
}

func handler(c *gee.Context) {
	html := "<h2>hello</h2>"
	c.HTML(200, html)
}

func counter(c *gee.Context) {
	c.Data(200, []byte("fadsfsa"))
}
