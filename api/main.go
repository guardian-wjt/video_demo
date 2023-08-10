package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	// 初始化 httprouter
	router := httprouter.New()

	router.POST("/user", CreateUser)

	return router
}

func main() { //放一些简单的定义性的东西(基于http分层原理)
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
