package main

func addRoutes(router *Router) {
	router.GET("/ping", func(req *Request) Response {
		return Response {
			Status: 200,
			Body: "pong",
		}
	})

	router.GET("/hello", func(req *Request) Response {
		return Response {
			Status: 200,
			Body: "hello from server 👋",
		}
	})

	router.POST("/echo", func(req *Request) Response {
		return Response {
			Status: 200,
			Body: string(req.Body),
		}
	})
}