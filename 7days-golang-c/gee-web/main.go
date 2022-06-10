package main

import (
	"example/gee"
	"net/http"
)

func main() {
	engine := gee.NewEngine()
	engine.Get("/", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	engine.Get("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	engine.Get("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	engine.Get("/assets/*filepath", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})
	engine.Run(":9999")
}
