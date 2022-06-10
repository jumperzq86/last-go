package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
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

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) getRoute(method string, part string) (*node, map[string]string) {

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	parts := parsePattern(part)
	n := root.search(parts, 0)
	if n != nil {
		pps := parsePattern(n.pattern)
		params := make(map[string]string)
		for index, pp := range pps {
			if pp[0] == ':' {
				params[pp[1:]] = parts[index]
			}
			if pp[0] == '*' && len(pp) > 1 {
				params[pp[1:]] = strings.Join(parts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *Router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		f := r.handlers[key]
		f(c)
	} else {
		c.String(http.StatusNotFound, "404 not found: %s\n", c.Path)
	}
}
