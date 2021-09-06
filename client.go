package main

// import (
// 	"fmt"
// 	lib "main/lib/net"
// 	"main/proto/main/pb"
// )

// func main() {
// 	router := lib.CreateRouter()
// 	router.RegisterHandler(int32(pb.PACKET_CMD_LOGIN_RES), func(msg, _ interface{}) interface{} {
// 		res := msg.(*pb.LoginRes)
// 		fmt.Println(res.Player.Name, res.Player.Gold, res.Player.Lv)
// 		return nil
// 	})
// 	convert := lib.NewProtoDataConvert()
// 	convert.RegisterProto(pb.PACKET_CMD_LOGIN_RES, func() interface{} { return &pb.LoginRes{} })
// 	client := lib.NewTcpClient("127.0.0.1", "8101", router, convert)
// 	client.Start()
// 	req := &pb.LoginReq{Base: &pb.BaseMessage{}}
// 	// cmd := pb.PACKET_CMD_EMPTY
// 	req.Base.Cmd = pb.PACKET_CMD_LOGIN
// 	req.Account = "test"

// 	client.Send(req)

// 	running := make(chan bool)
// 	<-running
// 	close(running)
// }
