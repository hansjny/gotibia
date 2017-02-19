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
//0-19: ??
//19-22: Acc nr
//23: pwlen
//24: wat?
//25: pw-pwlen

func handleRequest(con net.Conn) {
	buf := make([]byte, 1024)

	reqLen, err := con.Read(buf)

	if err != nil {
		fmt.Println("Error reading: ", err.Error());
	}
	var recvstr string = string(buf[:reqLen])
	var pwlen uint8 = uint8(buf[23])
	var acc_num uint32 = binary.LittleEndian.Uint32(buf[19:23])
	var pass = string(buf[25:25+pwlen])
	fmt.Println("Message received. Bytes: ", reqLen)
	fmt.Println("String: ", recvstr)
	fmt.Println("PW len: ", pwlen)
	fmt.Println("Acc: ", acc_num); 
	fmt.Println("Pass: ", pass)
	fmt.Println(buf);

	con.Write(
	con.Close();

}
