package main

func addRoutes(router *Router) {
	router.GET("/ping", func(req *Request) string {
		return string(req.Body)
	})

	router.GET("/hello", func(req *Request) string {
		return "hello from server 👋"
	})

	router.POST("/echo", func(req *Request) string {
		return string(req.Body)
	})
}