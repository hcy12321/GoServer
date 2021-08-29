package logic

import (
	"fmt"
	"main/entity"
	"main/proto/main/pb"
)

func Login(msg, _ interface{}) interface{} {
	req := msg.(*pb.LoginReq)
	res := &pb.LoginRes{Base: &pb.BaseMessage{}}
	player := &entity.PlayerEntity{}
	player.Load(req.Account)
	fmt.Printf("%v", *player)
	player.Name = req.Account
	player.Save()
	res.Base.Cmd = pb.PACKET_CMD_LOGIN_RES
	res.Player = &pb.Player{}
	res.Player.Name = player.Name
	res.Player.Gold = int32(player.Gold)
	res.Player.Lv = int32(player.Lv)

	return res
}
