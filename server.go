package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3838"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 12)
	reqLen, err := conn.Read(buf)
	if reqLen == 0 || err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	pack := binary.BigEndian.Uint64(buf[2:10])
	pack += 3600

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, pack)

	rePack := []byte{}
	rePack = append(rePack, 0xff, 0xff)
	rePack = append(rePack, b...)
	rePack = append(rePack, 0xff, 0xff)

	conn.Write(rePack)
	conn.Close()
}
