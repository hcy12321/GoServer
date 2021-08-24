package logic

import "main/proto/main/pb"

func Login(msg, _ interface{}) interface{} {
	req := msg.(*pb.LoginReq)
	res := &pb.LoginRes{Base: &pb.BaseMessage{}}

	res.Base.Cmd = pb.PACKET_CMD_LOGIN_RES
	res.Player = &pb.Player{}
	res.Player.Name = req.Account
	res.Player.Gold = 0
	res.Player.Lv = 1
	return res
}
