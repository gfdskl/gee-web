package gee

import (
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandleFunc),
	}
}

func (engine *Engine) addRouter(method string, pattern string, handle HandleFunc) {
	key := method + "-" + pattern
	engine.router[key] = handle
}

func (engine *Engine) GET(pattern string, handle HandleFunc) {
	engine.addRouter("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle HandleFunc) {
	engine.addRouter("POST", pattern, handle)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handle, ok := engine.router[key]; ok {
		handle(w, r)
	} else {
		http.Error(w, "not found", 404)
	}
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}
