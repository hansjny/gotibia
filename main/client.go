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

func (c *Client) sendDisconnectError(err string) {
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

/*Order of operations is important, as it follows
the Tibia protocol for receiving data */
func (c *Client)loginProtocol() account.Account {
	m := c.inBuf;
	m.skipBytes(5)

	//Check protocol
	proto := m.readUint16()
	if proto != 730 {
		c.sendDisconnectError(ERROR_PROTO)
		return account.Account{}
	}

	m.skipBytes(12)
	accNum := strconv.Itoa(int(m.readUint32()))
	pwlen := m.readUint16()
	password := m.readStr(int(pwlen)) 

	//Verify non empty acc or password
	if accNum == "0"  || pwlen == 0 {
		c.sendDisconnectError(ERROR_EMPTY_PW)
		return account.Account{}
	}

	id := account.RequestAccount(accNum, password)
	if id == 0 {
		c.sendDisconnectError(ERROR_WRONG_PW)
		return account.Account{}
	}
	c.sendDisconnectError(ERROR_UNDER_DEV)
	return account.Account{}
}


