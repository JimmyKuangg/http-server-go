package main

type HandlerFunc func(req *Request) string

type Router struct {
	routes map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.routes["GET "+path] = handler
}

func (r *Router) Handle(req *Request) string {
	key := req.Method + " " + req.Path

	if handler, ok := r.routes[key]; ok {
		return handler(req)
	}

	return "404 Not Found"
}