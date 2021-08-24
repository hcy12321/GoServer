package lib

import (
	"fmt"
	"net"
)

type TcpClient struct {
	host      string
	port      string
	conn      *CommonConn
	router    *Router
	converter DataConvertInterface
}

func NewTcpClient(host, port string, router *Router, converter DataConvertInterface) (c *TcpClient) {
	c = &TcpClient{}
	c.host = host
	c.port = port
	c.router = router
	c.converter = converter
	return
}

func (c *TcpClient) Start() {
	nativeConn, error := net.Dial("tcp", c.host+":"+c.port)
	if error != nil {
		fmt.Println("error dialing", error.Error())
		return
	}
	c.conn = &CommonConn{conn: &nativeConn, convert: c.converter, router: c.router}
	c.conn.Read()
}

func (c *TcpClient) Send(i interface{}) {
	c.conn.Send(i)
}
