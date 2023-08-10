package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// 存放逻辑处理

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}
