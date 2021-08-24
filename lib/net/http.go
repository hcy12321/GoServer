package lib

import (
	"fmt"
	"net/http"
)

type HttpServer struct {
	router *Router
}

func (server *HttpServer) SetRouter(router *Router) {
	server.router = router
}

func (server *HttpServer) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func (server *HttpServer) Create(addr string) {
	fmt.Println("Http Listen on", addr)
	http.HandleFunc("/", server.Handler)
	http.ListenAndServe(addr, nil)
}
