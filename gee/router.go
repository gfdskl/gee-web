package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	root   map[string]*node
	router map[string]HandleFunc
}

func newRouter() *Router {
	return &Router{
		root:   make(map[string]*node),
		router: make(map[string]HandleFunc),
	}
}

func splitParts(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) addRouter(method string, pattern string, handle HandleFunc) {
	key := method + "-" + pattern
	r.router[key] = handle
	if r.root[method] == nil {
		r.root[method] = newNode()
	}
	p := r.root[method]
	parts := splitParts(pattern)
	p.insert(pattern, parts, 0)
}

func (r *Router) getRouter(method string, pattern string) (HandleFunc, map[string]string) {
	if r.root[method] == nil {
		return nil, nil
	}
	p := r.root[method]
	parts := splitParts(pattern)
	res := p.search(parts, 0)
	if res == nil {
		return nil, nil
	}
	parms := make(map[string]string)
	matched := splitParts(res.pattern)
	handle := r.router[method+"-"+res.pattern]
	for i, match := range matched {
		if match[0] == ':' {
			parms[match[1:]] = parts[i]
		} else if match[0] == '*' {
			parms[match[1:]] = strings.Join(parts[i:], "/")
		}
	}
	return handle, parms
}

func (r *Router) handle(c *Context) {
	handler, params := r.getRouter(c.Method, c.Path)
	if handler == nil {
		http.Error(c.Writer, "not found", 500)
		return
	}
	c.Params = params
	handler(c)
}
