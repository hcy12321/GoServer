package lib

import (
	"fmt"
	"net"
)

type TcpServer struct {
	host      string
	port      string
	connMap   map[string]*TcpConn
	router    *Router
	converter DataConvertInterface
}

func NewTcpServer(host, port string, router *Router, converter DataConvertInterface) *TcpServer {
	srv := &TcpServer{}
	srv.host = host
	srv.port = port
	srv.router = router
	srv.converter = converter
	srv.connMap = make(map[string]*TcpConn)
	return srv
}

func (srv *TcpServer) Start() {
	ln, _ := net.Listen("tcp", srv.host+":"+srv.port)
	fmt.Printf("tcp listen on %s:%s", srv.host, srv.port)

	go func() {
		for {
			naticeConn, _ := ln.Accept()
			tcpConn := NewTcpConn(srv, &naticeConn, srv.router, srv.converter)
			addr := (*(tcpConn.conn.conn)).RemoteAddr().String()
			srv.connMap[addr] = tcpConn
			tcpConn.Read()
			fmt.Println("a client connected :" + addr)
		}
	}()
}

func (srv *TcpServer) Broadcast(i interface{}) {
	for _, v := range srv.connMap {
		v.Send(i)
	}
}
