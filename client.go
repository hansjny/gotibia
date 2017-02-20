package main

import (
	"fmt"
	"net"
	//"encoding/hex"
)
type Client struct {
	socket net.Conn
	inBuf Msg
	outBuf *Msg
}

func NewClient(s net.Conn) *Client {
	c := Client{socket: s}
	c.outBuf = NewMsg(2);
	return &c
}

func (c *Client)loginProtocol() {
	m := c.inBuf;
	m.skipBytes(5)
	proto := m.readUint16()
	fmt.Printf("Protocol: %d\n", proto)
	inBuf := "Server corrently in developement."
	c.sendError(inBuf)
}

func (c *Client) sendError(err string) {
	var byteval uint8 = TIBIA_LOGIN_ERROR;
	var msglen uint16 = uint16(len(err));
	c.outBuf.addUint8(byteval)
	c.outBuf.addUint16(msglen)
	c.outBuf.addString(err)
	c.sendMessage()
}

func (c *Client) remoteAddr() string {
	return c.socket.RemoteAddr().String()
}

func (c *Client) addHeader() {
	c.outBuf.pos = 0;	
	fmt.Println("BUFFER LENGTH: ", len(c.outBuf.data))
	//c.outBuf.addUint16(uint16(len(err) + 3))
	
}

func (c *Client) sendMessage() {
	c.addHeader();
	c.socket.Write(c.outBuf.data)
	c.outBuf = NewMsg(2)
}

func (c *Client) receive() {
	c.inBuf.newMessage(c.socket);
}
