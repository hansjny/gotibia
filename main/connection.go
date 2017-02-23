package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	account "github.com/hansjny/tibia71go/account"
	"net"
	"os"
)

const (
	CONN_HOST = ""
	CONN_PORT = "7171"
	CONN_TYPE = "tcp"
)

func main() {
	//	Connection socket listening
	l, err := net.Listen("tcp", HOST_IP+":"+HOST_PORT)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + HOST_IP + ":" + HOST_PORT)

	db, dberr := sql.Open("mysql", MYSQL_USER + ":"+ MYSQL_PW +
						  "@tcp("+ MYSQL_IP+ ":"+ MYSQL_PORT +")/" + MYSQL_DB)
	if dberr != nil {
		fmt.Println("Error connecting to DB: ", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	for {
		fmt.Println("Waiting for connection...")
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
	account.RequestAccount("195176")
	//con.Close()
}
