package lib

import (
	"main/proto/main/pb"
	"unsafe"
)

type ProtoBaseInterface interface {
	GetBase() *(pb.BaseMessage)
}

type DataConvertInterface interface {
	Decode([]byte) interface{}
	Encode(interface{}) []byte
}

type BaseDataConvert struct {
}

func (convert *BaseDataConvert) Init() {

}

func (convert *BaseDataConvert) Decode(b []byte) interface{} {
	str := *(*string)(unsafe.Pointer(&b))
	// str := string(b)
	return str
}

func (convert *BaseDataConvert) Encode(data interface{}) []byte {
	str := data.(string)
	// return []byte(str)
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
