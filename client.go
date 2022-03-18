package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:3838")

	i := time.Now().Unix()

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))

	pack := []byte{}
	pack = append(pack, 0xff, 0xff)
	pack = append(pack, b...)
	pack = append(pack, 0xff, 0xff)

	n, err := conn.Write([]byte(pack))
	if n == 0 || err != nil {
		return
	}

	repack := make([]byte, 12)
	message, _ := bufio.NewReader(conn).Read(repack)
	if message == 0 {
		return
	}
	fmt.Println(binary.LittleEndian.Uint64(repack[2:10]))

}
