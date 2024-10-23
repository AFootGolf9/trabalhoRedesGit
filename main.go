package main

import (
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "144.22.201.166:9000")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// spam(conn)
	basic(conn)
}

func basic(conn net.Conn) {
	conn.Write([]byte("ASK IP"))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	println(string(buf[:n]))

	conn.Write([]byte("REG USER 1436 parei"))

	n, _ = conn.Read(buf)

	temp := string(buf[:n])
	println(temp)

	if temp != "RESPONSE 200 " {
		panic("Algo errado")
	}

	conn.Write([]byte("ASK SECRET 1436 parei"))

	n, _ = conn.Read(buf)

	println(string(buf[:n]))
}

func spam(conn net.Conn) {

	for i := 1400; i <= 2000; i++ {
		msg := "REG USER " + strconv.Itoa(i) + " SPAM"
		println(msg)
		conn.Write([]byte(msg))

		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)

		temp := string(buf[:n])
		println(temp)

		if temp != "RESPONSE 200 " {
			// panic("Algo errado")
		}
	}
}
