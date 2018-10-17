package advanced_examples

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("server", string(data[:n]))

		c.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go requestHandler(conn)
	}
}
