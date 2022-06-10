package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *Router
}

func NewEngine() *Engine {
	return &Engine{router: NewRouter()}
}

func (e *Engine) addRouter(method string, pattern string, f HandlerFunc) {
	e.router.addRoute(method, pattern, f)
}

func (e *Engine) Get(pattern string, f HandlerFunc) {
	e.addRouter("GET", pattern, f)
}

func (e *Engine) Post(pattern string, f HandlerFunc) {
	e.addRouter("POST", pattern, f)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(r, w)
	e.router.handle(c)
}

func (e *Engine) Run(port string) {
	http.ListenAndServe(port, e)
}
