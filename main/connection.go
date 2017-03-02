package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	color "github.com/fatih/color"
	account "github.com/hansjny/tibia71go/account"
	"net"
	"os"
)

type World struct {
	Name string
	Id int
	Port int
	Ip string
}

var G_DB *sql.DB
var WORLDLIST []World

/* Server initialization procedure */
func main() {

	fmt.Println("::::::::::::::::::: GoTibia 7.3 ::::::::::::::::::")
	loadPrint("MySql connection", connectDb())
	account.G_DB = G_DB
	loadPrint("Loading worlds",  loadWorlds())
	fmt.Printf("%-*s", 40, ":: Setting up TCP socket")
	err := connectTcpSocket()
	if (err != nil) {
		fmt.Printf("%*s", 20, color.RedString("[ Failed ]\n"))
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func loadPrint(msg string, err error) {
	fmt.Printf("%-*s", 40, ":: "+msg)
	if (err != nil) {
		fmt.Printf("%*s", 20, color.RedString("[ Failed ]\n"))
		fmt.Println("Error message: ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%*s", 20, color.GreenString("[ OK ]\n"))
}

func loadWorlds() error {
	rows,  err := G_DB.Query("SELECT * FROM worlds")

	if err != nil {
		fmt.Println("Something went wrong,  loadWorlds()")
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var wname string
		var wid int
		var wip string
		var wport int 
		err := rows.Scan(&wid, &wname, &wip,  &wport)
		if err != nil {
			fmt.Println("Something went wrong,  loadWorlds()")
			return err
		}
		w := World{Name: wname,  Id: wid,  Port: wport,  Ip: wip}
		WORLDLIST = append(WORLDLIST,  w)
	}
	return err
}
func connectDb() error {
	db, dberr := sql.Open("mysql", MYSQL_USER + ":"+ MYSQL_PW +
	"@tcp("+ MYSQL_IP+ ":"+ MYSQL_PORT +")/" + MYSQL_DB)
	G_DB = db

	rows, err := db.Query("SELECT * FROM accounts")
	if (err != nil) {
		return err
	}
	_ = rows

	return dberr
}

func connectTcpSocket() error {
	//	Connection socket listening
	l, err := net.Listen("tcp", HOST_IP+":"+HOST_PORT)
	if err != nil {
		return err;
	}
	defer l.Close()

	fmt.Printf("%*s", 20, color.GreenString("[ OK ]\n"))
	fmt.Print("\n\nListening on IP: ")
		if HOST_IP == "" {
			fmt.Print("all addresses")
		} else {
			fmt.Print(HOST_IP)
		}

	fmt.Println(", port: " + HOST_PORT)
	color.Green("[Server running!]")

	for {
		con, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handleRequest(con)
	}
	return err;
}

func handleRequest(con net.Conn) {
	client := NewClient(con)
	fmt.Printf("Remote connection from %s accepted\n", client.remoteAddr())
	client.receive()
	msg := client.inBuf;
	msglen := msg.readUint16()
	msgtype := msg.readUint8()
	if (msgtype == 0x1) {
		msg.skipBytes(2)
		fmt.Printf("Msglen: %d:, msg type: %d\n", msglen, msgtype)
		if client.loginProtocol() == nil {
			fmt.Printf("Wrong account number or password.")
		}
	} else {

		fmt.Println("Dumping packet")
		client.dumpOutputPacket()

	}
	//account.RequestAccount("195176", "hello")
	//con.Close()
}
