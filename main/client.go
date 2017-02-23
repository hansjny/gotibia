package main

import (
	"fmt"
	"net"
	"strconv"
	"encoding/hex"
	account "github.com/hansjny/tibia71go/account"
)
type Client struct {
	socket net.Conn
	inBuf *IncomingMessage
	outBuf *OutgoingMessage
}

func NewClient(s net.Conn) *Client {
	c := Client{socket: s}
	c.outBuf = NewOutgoingMessage(2);
	c.inBuf = NewIncomingMessage(1024)
	return &c
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
	a, b := c.outBuf.Uint16ToUint8(uint16(len(c.outBuf.data)-2))
	c.outBuf.data[0] = b
	c.outBuf.data[1] = a 
} 

func (c *Client) dumpPacket() {
	fmt.Println(hex.Dump(c.outBuf.data))
}
func (c *Client) sendMessage() {
	c.addHeader();
	c.dumpPacket();
	c.socket.Write(c.outBuf.data)
	c.outBuf = NewOutgoingMessage(2)
}

func (c *Client) receive() {
	c.inBuf.newMessage(c.socket);
}

func (c *Client)loginProtocol() account.Account {
	m := c.inBuf;
	m.skipBytes(5)

	//Check protocol
	proto := m.readUint16()
	if proto != 730 {
		errMsg := "Tibia client 7.3 required to play."
		c.sendError(errMsg)
		return account.Account{}
	}

	m.skipBytes(12)
	accNum := strconv.Itoa(int(m.readUint32()))
	account.RequestAccount(accNum)
	pwlen := m.readUint8()
	password := m.readStr(int(pwlen)) 

	//Verify non empty acc or password
	if accNum == "0"  || pwlen == 0 {
		errMsg := "Account number or password can't be empty."
		c.sendError(errMsg)
		return account.Account{}
	}
	_ = password
	errMsg := "Server currently in developement."	
	//inBuf := "Please enter a valid account number and password."	
	c.sendError(errMsg)
	return account.Account{}
}


