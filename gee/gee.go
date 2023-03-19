package gee

import (
	"net/http"
)

type HandleFunc func(context *Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) addRouter(method string, pattern string, handle HandleFunc) {
	engine.router.addRouter(method, pattern, handle)
}

func (engine *Engine) GET(pattern string, handle HandleFunc) {
	engine.addRouter("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle HandleFunc) {
	engine.addRouter("POST", pattern, handle)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	engine.router.handle(context)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}
