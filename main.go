package main

import (
	lib "main/lib/net"
	"main/logic"
	"main/proto/main/pb"
)

func main() {
	httpSrv := &lib.HttpServer{}
	router := lib.CreateRouter()
	router.RegisterHandler(int32(pb.PACKET_CMD_LOGIN), logic.Login)
	httpSrv.SetRouter(router)
	go func() {
		httpSrv.Create("127.0.0.1:8001")
	}()
	convert := lib.NewProtoDataConvert()
	convert.RegisterProto(int32(pb.PACKET_CMD_LOGIN), &pb.LoginReq{})
	tcpSrv := lib.NewTcpServer("127.0.0.1", "8101", router, convert)
	tcpSrv.Start()

	running := make(chan bool)
	<-running
	close(running)
}
