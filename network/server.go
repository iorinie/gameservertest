package network

import (
	"fmt"
	"net"
)

func StartServer() {
	fmt.Println("SERVER|listening...")

	listener, err := net.Listen("tcp", ":10010")
	if err != nil {
		fmt.Printf("SERVER|tcp listen to port fail, err = %s\n", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("SERVER|listener accept fail, err = %s\n", err)
			continue
		}

		fmt.Printf("SERVER|a client is connected, remote = %v\n", conn.RemoteAddr())

		go handleCConn(conn)
	}
}

func handleCConn(conn net.Conn) {
	for {
		recvData := make([]byte, 32)
		n, err := conn.Read(recvData)
		if err != nil {
			fmt.Printf("SERVER|read from client fail, err = %s\n", err)
			conn.Close()
			break
		}

		fmt.Printf("SERVER|receive data from client: %v\n", string(recvData[:n]))

		conn.Write([]byte("hi, i am server."))
	}
}