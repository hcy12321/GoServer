package lib

import (
	"bufio"
	"encoding/binary"
	"net"
)

type CommonConn struct {
	conn    *net.Conn
	router  *Router
	convert DataConvertInterface
}

func (c *CommonConn) Read() {
	reader := bufio.NewReader(*(c.conn))
	scanner := bufio.NewScanner(reader)

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		la := uint32(len(data))
		if la < 4 {
			return 0, nil, nil
		}

		lc := binary.LittleEndian.Uint32(data[:4]) + 4
		if la >= lc {
			return int(lc), data[4:lc], nil
		}
		if atEOF {
			return 0, data, bufio.ErrFinalToken
		}
		return
	}

	scanner.Split(split)

	go func() {
		for scanner.Scan() {
			msg := c.convert.Decode(scanner.Bytes())
			// v := reflect.ValueOf(msg)
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
			packet, ok := msg.(ProtoBaseInterface)
			if !ok {
				return
			}

			base := packet.GetBase()
			if base == nil {
				return
			}
			cmd := base.Cmd
			msgId := int32(cmd)
			res := c.router.Process(msgId, msg, c.conn)
			if res != nil {
				c.Send(res)
			}
		}
	}()
}

func (c *CommonConn) Send(i interface{}) {
	b := c.convert.Encode(i)
	c.Write(b)
}

func (c *CommonConn) Write(b []byte) {
	if len(b) == 0 {
		return
	}

	a := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, uint32(len(b)))
	(*(c.conn)).Write(a)
	(*(c.conn)).Write(b)
}
