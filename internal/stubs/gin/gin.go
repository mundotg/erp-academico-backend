package gin

import "net/http"

type H map[string]interface{}
type HandlerFunc func(*Context)
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
}

func (c *Context) JSON(code int, obj interface{}) {
	if c.Writer != nil {
		c.Writer.WriteHeader(code)
	}
}
func (c *Context) ShouldBindJSON(obj interface{}) error { return nil }
func (c *Context) Param(key string) string {
	if c.Params == nil {
		return ""
	}
	return c.Params[key]
}

type RouterGroup struct{}
type Engine struct{ RouterGroup }

func Default() *Engine                                           { return &Engine{} }
func (e *Engine) Run(...string) error                            { return nil }
func (g *RouterGroup) Group(string, ...HandlerFunc) *RouterGroup { return &RouterGroup{} }
func (g *RouterGroup) GET(string, ...HandlerFunc)                {}
func (g *RouterGroup) POST(string, ...HandlerFunc)               {}
