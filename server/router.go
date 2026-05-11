package main

type HandlerFunc func(req *Request) string

type Router struct {
	routes map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (r *Router) addRoute(method, path string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}

	r.routes[method][path] = handler
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.addRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.addRoute("POST", path, handler)
}

func (r *Router) Handle(req *Request) string {
	if r.routes[req.Method] != nil {
		if handler, ok := r.routes[req.Method][req.Path]; ok {
			return handler(req)
		}
	}

	return "404 Not Found"
}