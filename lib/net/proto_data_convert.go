package lib

import (
	"encoding/binary"
	"log"
	"reflect"

	"google.golang.org/protobuf/proto"
)

type ProtoDataConvert struct {
	protoMap map[int32]reflect.Type
}

func NewProtoDataConvert() *ProtoDataConvert {
	convert := &ProtoDataConvert{}
	convert.protoMap = make(map[int32]reflect.Type)
	return convert
}

func (cdc *ProtoDataConvert) RegisterProto(msgId int32, msg interface{}) {
	msgType := reflect.TypeOf(msg.(proto.Message))
	cdc.protoMap[msgId] = msgType
}

func (convert *ProtoDataConvert) Decode(b []byte) (msg interface{}) {
	if len(b) < 4 {
		return
	}

	msgId := binary.LittleEndian.Uint32(b[:4])
	if msgType, ok := convert.protoMap[int32(msgId)]; ok {
		msg = reflect.New(msgType.Elem()).Interface()
		err := proto.Unmarshal(b[4:], msg.(proto.Message))
		if err != nil {
			log.Fatal(err)
			return
		}

		return
	}
	return nil
}

func (convert *ProtoDataConvert) Encode(data interface{}) (b []byte) {
	// v := reflect.ValueOf(data)
	// base := reflect.Indirect(v).FieldByName("Base")
	// if !base.IsValid() {
	// 	return
	// }
	// cmd := reflect.Indirect(base).FieldByName("Cmd")
	// if !cmd.IsValid() {
	// 	return
	// }
	// if cmd.Type().Kind() != reflect.Int32 {
	// 	return
	// }

	msg, ok := data.(ProtoBaseInterface)
	if !ok {
		return
	}
	base := msg.GetBase()
	if base == nil {
		return
	}
	cmd := base.Cmd
	h := make([]byte, 4)
	msgId := uint32(cmd)
	binary.LittleEndian.PutUint32(h, msgId)

	t, err := proto.Marshal(data.(proto.Message))

	if err != nil {
		log.Fatal(err)
		return
	}

	b = append(h, t...)
	return
}
