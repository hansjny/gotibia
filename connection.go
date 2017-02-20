package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = ""
	CONN_PORT = "7171"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+ CONN_PORT)

	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}


	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		con, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(con)
	}
}

func handleRequest(con net.Conn) {
	client := NewClient(con)
	fmt.Printf("Remote connection from %s accepted\n", client.remoteAddr())
	client.receive()
	client.loginProtocol()


	/*
	var recvstr string = string(buf[:reqLen])
	var pwlen uint8 = uint8(buf[23])
	var acc_num uint32 = binary.LittleEndian.Uint32(buf[19:23])
	var pass = string(buf[25:25+pwlen]) fmt.Println("Message received. Bytes: ", reqLen)
	fmt.Println("String: ", recvstr)
	fmt.Println("PW len: ", pwlen)
	fmt.Println("Acc: ", acc_num)
	fmt.Println("Pass: ", pass)
	fmt.Println(hex.Dump(buf))
	con.Write([]byte("AAAAAAAAAAB")) */
	con.Close()
}
