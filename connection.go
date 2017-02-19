package main

import (
	"fmt"
	"net"
	"os"
	"encoding/binary"
)

const (
	CONN_HOST = "192.168.1.9"
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
		con, err := l.Accept();
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
	go handleRequest(con)
	}
}

//Protokoll
//0-20: ??
//20-23: Acc nr
//24: pwlen
//25: wat?
//26: pw-pwlen

func handleRequest(con net.Conn) {
	buf := make([]byte, 1024)

	reqLen, err := con.Read(buf)

	if err != nil {
		fmt.Println("Error reading: ", err.Error());
	}

	fmt.Println("Message received. Bytes: ", reqLen)
	fmt.Println(binary.LittleEndian.Uint32(buf[19:23]))
	fmt.Println(string(buf[:reqLen]))
	fmt.Println(buf);

	con.Close();

}
