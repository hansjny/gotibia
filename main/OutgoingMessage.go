// Package messaging provides messaging functions
package main

import (
//	"encoding/binary"
	"fmt"
//	"net"
)
type OutgoingMessage struct {
	pos int
	data []byte
}

func NewOutgoingMessage(dataAmount int) *OutgoingMessage {
	m := &OutgoingMessage{pos: dataAmount}
	m.data = make([]byte, dataAmount)
	return m
}


func (m* OutgoingMessage) Uint16ToUint8(val uint16) (uint8, uint8) {
	return uint8(val >> 8), uint8(val)
}

func (m *OutgoingMessage) addUint16(val uint16) {
	a, b := m.Uint16ToUint8(val)
	m.data = append(m.data, b, a)
	fmt.Printf("Packet type: 0x%x%x\n", m.data[m.pos],m.data[m.pos+1])
	m.pos += 2
}

func (m *OutgoingMessage) addUint8(val uint8) {
	m.data = append(m.data, val)
	fmt.Printf("Packet type: 0x%02x\n", m.data[m.pos])
	m.pos += 1
}

func (m *OutgoingMessage) msgHex() {


}

func (m *OutgoingMessage) addString(str string) {
	app := []byte(str)
	m.data = append(m.data, app...)
	m.pos += len(str)
}
