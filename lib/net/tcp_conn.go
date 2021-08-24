package lib

import "net"

type TcpConn struct {
	tcpServer *TcpServer
	conn      *CommonConn
}

func NewTcpConn(tcpServer *TcpServer, nativeConn *net.Conn, router *Router, convert DataConvertInterface) *TcpConn {
	c := &TcpConn{}
	c.tcpServer = tcpServer
	c.conn = &CommonConn{nativeConn, router, convert}
	return c
}

func (c *TcpConn) Read() {
	c.conn.Read()
}

func (c *TcpConn) Send(i interface{}) {
	c.conn.Send(i)
}
