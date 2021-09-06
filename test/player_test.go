package test

import (
	"fmt"
	lib "main/lib/net"
	"main/logic"
	"main/proto/main/pb"
	"os"
	"path"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogin(t *testing.T) {

	Convey("test login", t, func() {
		exePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println("load exePath err")
		}

		configDir := path.Join(exePath, "config")
		configPath := path.Join(configDir, "base.json")
		_, sErr := os.Stat(configPath)
		if sErr != nil {
			fmt.Println(sErr)
			if os.IsNotExist(sErr) {
				os.Mkdir(configDir, 0666)
				convert := &lib.BaseDataConvert{}
				os.WriteFile(configPath, convert.Encode("{    \"defaultRedisConfig\": {        \"host\": \"127.0.0.1\",        \"port\": 6379,        \"pwd\": \"hcy!@#123\",        \"db\": 0    }}"), 0666)
			}
		}

		req := &pb.LoginReq{Base: &pb.BaseMessage{Cmd: pb.PACKET_CMD_LOGIN}, Account: "tete"}
		res := logic.Login(req, nil)

		Convey("res is not nil", func() {
			So(res, ShouldNotBeNil)
		})
	})
}
