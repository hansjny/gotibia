package main

import (
	"fmt"
	"net"
	"strconv"
	"encoding/hex"
	"bytes"
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
	c.outBuf.addUint8(byteval)
	c.outBuf.addString(err)
	c.sendMessage()
}

func (c *Client) sendCharPick(acc *account.Account) {
	var buffer bytes.Buffer
	buffer.WriteByte(0x31) //Motd ID
	buffer.WriteByte(0xA) //Newline
	buffer.WriteString("Welcome to Netherfall!")

	//Accepted PW
	c.outBuf.addUint8(0x14) 
	//MOTD
	c.outBuf.addString(buffer.String()) //String

	//Charlist
	c.outBuf.addUint8(0x64) 

	c.outBuf.addUint8(uint8(len(acc.Chars))) //Charlist size
	for _,  char := range acc.Chars {
		world := WORLDLIST[char.WorldId - 1]
		c.outBuf.addString(char.Name) //Name
		c.outBuf.addString(world.Name)
		ip := net.ParseIP(world.Ip)
		c.outBuf.addUint8(ip[12]) //ip
		c.outBuf.addUint8(ip[13]) //ip
		c.outBuf.addUint8(ip[14]) //ip
		c.outBuf.addUint8(ip[15]) //ip
		c.outBuf.addUint16(uint16(world.Port)) //Port
	}

	//Add premium days
	c.outBuf.addUint16(uint16(acc.Prem)) 
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

func (c *Client) dumpOutputPacket() {
	fmt.Println(hex.Dump(c.outBuf.data))
}

func (c *Client) dumpIncomingPacket() {
	fmt.Println(hex.Dump(c.inBuf.data))
}
func (c *Client) sendMessage() {
	c.addHeader();
	c.dumpOutputPacket();
	c.socket.Write(c.outBuf.data)
	c.outBuf = NewOutgoingMessage(2)
}

func (c *Client) receive() {
	c.inBuf.newMessage(c.socket);
}

/*Order of operations is important, as it follows
the Tibia protocol for receiving data */

func (c *Client)loginProtocol() *account.Account {
	m := c.inBuf;

	//Check protocol
	proto := m.readUint16()
	if proto != TIBIA_ALLOWED_PROTOCOL {
		c.sendDisconnectError(ERROR_PROTO)
		return nil
	}

	m.skipBytes(12)
	accNum := strconv.Itoa(int(m.readUint32()))
	pwlen := m.readUint16()
	password := m.readStr(int(pwlen)) 

	//Verify non empty acc or password
	if accNum == "0"  || pwlen == 0 {
		c.sendDisconnectError(ERROR_EMPTY_PW)
		return nil
	}

	acc := account.RequestAccount(accNum, password)
	if acc == nil {
		c.sendDisconnectError(ERROR_WRONG_PW)
		return nil
	}
	//c.sendDisconnectError(ERROR_UNDER_DEV)
	c.sendCharPick(acc)
	return acc
}


