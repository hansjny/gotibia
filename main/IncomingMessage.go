// Package messaging provides messaging functions
package main

import (
	"encoding/binary"
	"fmt"
	"net"
)
type IncomingMessage struct {
	pos int
	data []byte
}

func NewIncomingMessage(dataAmount int) *IncomingMessage {
	m := &IncomingMessage{pos: 0}
	m.data = make([]byte, 1024)
	return m
}

func (m *IncomingMessage) newMessage(con net.Conn) {
	bytesread, err := con.Read(m.data)
	_ = bytesread
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	m.pos = 0
}

func (m *IncomingMessage) readStr(strlen int) string {
	var data string = string(m.data[m.pos:m.pos+strlen+1])
	m.pos += strlen
	return data
}
func (m *IncomingMessage) readUint16() uint16 {
	var data uint16 = binary.LittleEndian.Uint16(m.data[m.pos:m.pos+2])
	m.pos += 2
	return data
}

func (m *IncomingMessage) readUint32() uint32 {
	var data uint32 = binary.LittleEndian.Uint32(m.data[m.pos:m.pos+4])
	m.pos += 4
	return data
}
func (m *IncomingMessage) readUint8() uint8 {
	var data uint8 = uint8(m.data[m.pos])
	m.pos += 1
	return data
}

func (m *IncomingMessage) skipBytes(amount int) {
	m.pos += amount
}


