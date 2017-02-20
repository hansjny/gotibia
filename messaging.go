// Package messaging provides messaging functions
package main

import (
	"encoding/binary"
	"fmt"
	"net"
)
type Msg struct {
	pos int
	data []byte
}

func NewMsg(p int) *Msg {
	m := &Msg{pos: p}
	m.data = make([]byte, 1024)
	return m
}
func (m *Msg) newMessage(con net.Conn) {
	bytesread, err := con.Read(m.data)
	_ = bytesread
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	m.pos = 0
}

func (m Msg) readUint16() uint16 {
	var data uint16 = binary.LittleEndian.Uint16(m.data[m.pos:m.pos+2])
	m.pos += 2
	return data
}

func (m *Msg) skipBytes(amount int) {
	m.pos += amount
}

func (m *Msg) addUint16(val uint16) {
	var a, b uint8 = uint8(val >> 8), uint8(val)
	m.data = append(m.data, b, a)
	fmt.Printf("Packet type: 0x%x%x\n", m.data[m.pos],m.data[m.pos+1])
	m.pos += 2
}

func (m *Msg) addUint8(val uint8) {
	m.data = append(m.data, val)
	fmt.Printf("Packet type: 0x%02x\n", m.data[m.pos])
	m.pos += 1
}

func (m *Msg) msgHex() {


}
func (m *Msg) addString(str string) {
	app := []byte(str)
	m.data = append(m.data, app...)
	m.pos += len(str)
}
