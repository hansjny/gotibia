// Package messaging provides messaging functions
package main

import (
	"encoding/binary"
	"fmt"
	"net"
)
type Msg struct {
	content []byte
	pos int
}

func newMessage(con net.Conn) *Msg {
	var msg Msg
	msg.content = make([]byte, 1024)
	bytesread, err := con.Read(msg.content)
	_ = bytesread
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}

	msg.pos = 0
	return &msg
}

func (m *Msg) readUint16() uint16 {
	var data uint16 = binary.LittleEndian.Uint16(m.content[m.pos:m.pos+2])
	m.pos += 2
	return data
}

func (m *Msg) skipBytes(amount int) {
	m.pos += amount
}


