package main

import (
	"fmt"
	"net"
)
type Client struct {
	socket net.Conn
	msg *Msg
}

func (c *Client)loginProtocol() {
	m := c.msg;
	m.skipBytes(5)
	proto := m.readUint16()
	fmt.Printf("Protocol: %d\n", proto)
}
