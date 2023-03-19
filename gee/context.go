package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/spf13/cobra"
)

type Context struct {
	// 基本信息
	Writer http.ResponseWriter
	Req    *http.Request
	// req信息
	Path   string
	Method string
	// writer信息
	StatusCode int
	// 动态路由参数
	Params map[string]string
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (context *Context) Param(key string) string {
	return context.Params[key]
}

func (context *Context) Query(key string) string {
	return context.Req.URL.Query().Get(key)
}

func (context *Context) PostForm(key string) string {
	return context.Req.FormValue(key)
}

func (context *Context) SetHeader(key string, value string) {
	context.Writer.Header().Set(key, value)
}

func (context *Context) SetStatusCode(code int) {
	context.StatusCode = code
	context.Writer.WriteHeader(code)
}

func (context *Context) String(code int, format string, values ...interface{}) {
	context.SetHeader("Content-Type", "text/plain")
	context.SetStatusCode(code)
	context.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (context *Context) JSON(code int, obj interface{}) {
	context.SetHeader("Content-Type", "application/json")
	context.SetStatusCode(code)
	encoder := json.NewEncoder(context.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Writer, err.Error(), 500)
	}
}

func (context *Context) Data(code int, data []byte) {
	context.SetStatusCode(code)
	context.Writer.Write(data)
}

func (context *Context) HTML(code int, html string) {
	context.SetHeader("Content-Type", "text/html")
	context.SetStatusCode(code)
	context.Writer.Write([]byte(html))
}
